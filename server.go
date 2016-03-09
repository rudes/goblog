package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
	"time"
)

const (
	STATIC_URL  = "/static/"
	STATIC_ROOT = "/home/rudes/go/src/github.com/rudes/OtherLetter/static/"
	TEMPLATES   = "/home/rudes/go/src/github.com/rudes/OtherLetter/templates"
)

var LoggedIn = false

type Post struct {
	ID, Title, Content, Date, Time string
}

type Context struct {
	Posts    []Post
	LoggedIn bool
	Static   string
}

func Home(w http.ResponseWriter, r *http.Request) {
	p, err := GetAllLetters()
	if err != nil {
		LogError(err)
		return
	}
	context := Context{
		Posts: p,
	}
	render(w, "index", context)
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	context := Context{}
	if r.Method == "GET" {
		render(w, "login", context)
	} else if r.Method == "POST" {
		name := r.FormValue("username")
		pass := r.FormValue("password")
		LogAnything("Attempting to Login: ", name)
		redirectTarget := "/"
		if name == "rudes" && pass == "demonking" {
			LoggedIn = true
			LogAnything("Login Successful")
			redirectTarget = "/new"
		}
		http.Redirect(w, r, redirectTarget, 302)
	}
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	context := Context{}
	LoggedIn = false
	LogAnything("Logged Out Successfully")
	render(w, "index", context)
}

func New(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		context := Context{}
		render(w, "new", context)
	} else if r.Method == "POST" {
		content := r.FormValue("content")
		title := r.FormValue("title")
		p := Post{
			Content: content,
			Title:   title,
		}
		HandleThisLetter(p)
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Path[len("/show/"):]
	if len(postID) != 0 {
		var p []Post
		post, err := GetThisLetter(postID)
		if err != nil {
			LogError(err)
			return
		}
		p = append(p, *post)
		context := Context{
			Posts: p,
		}
		render(w, "show", context)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Path[len("/edit/"):]
	if len(postID) != 0 {
		if r.Method == "GET" {
			var p []Post
			post, err := GetThisLetter(postID)
			if err != nil {
				LogError(err)
				return
			}
			p = append(p, *post)
			context := Context{
				Posts: p,
			}
			render(w, "edit", context)
		} else if r.Method == "POST" {
			var p Post
			p.ID = postID
			p.Title = r.FormValue("title")
			p.Content = r.FormValue("content")
			err := UpdateThisLetter(p)
			if err != nil {
				LogError(err)
				return
			}
		}
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Path[len("/delete/"):]
	if len(postID) != 0 {
		err := DeleteThisLetter(postID)
		if err != nil {
			LogError(err)
			return
		}
		fmt.Fprintf(w, "Deleted Post %s successfully", postID)
	}
}

func Static(w http.ResponseWriter, r *http.Request) {
	sf := r.URL.Path[len(STATIC_URL):]
	if len(sf) != 0 {
		f, err := http.Dir(STATIC_ROOT).Open(sf)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, r, sf, time.Now(), content)
			return
		}
	}
	http.NotFound(w, r)
}

func HandleLetters(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var p Post
	err := d.Decode(&p)
	if err != nil {
		LogError(err)
		return
	}
	HandleThisLetter(p)
}

func render(w http.ResponseWriter, tmpl string, context Context) {
	context.Static = STATIC_URL
	context.LoggedIn = LoggedIn
	tl := []string{TEMPLATES + "/base.tmpl", TEMPLATES + "/" + tmpl + ".tmpl"}
	t, err := template.ParseFiles(tl...)
	if err != nil {
		LogError(err)
		return
	}
	err = t.Execute(w, context)
	if err != nil {
		LogError(err)
		return
	}
}

func main() {
	http.HandleFunc("/api", HandleLetters)
	http.HandleFunc("/", Home)
	http.HandleFunc("/show/", Show)
	http.HandleFunc("/edit/", Edit)
	http.HandleFunc("/login/", LogIn)
	http.HandleFunc("/logout/", LogOut)
	http.HandleFunc("/delete/", Delete)
	http.HandleFunc(STATIC_URL, Static)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
