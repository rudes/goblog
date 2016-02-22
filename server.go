package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

const (
	STATIC_URL  = "/static/"
	STATIC_ROOT = "/go/src/github.com/rudes/OtherLetters/static/"
	TEMPLATES   = "/go/src/github.com/rudes/OtherLetters/templates"
)

// Structure for json data
type Post struct {
	ID, Title, Content, Date, Time string
}

type Context struct {
	Posts  []Post
	Static string
}

func Home(w http.ResponseWriter, r *http.Request) {
	p := GetAllLetters()
	context := Context{
		Posts: p,
	}
	render(w, "index", context)
}

// Handle incoming post data
// Submit data into sql
func HandleLetters(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var p Post
	err := d.Decode(&p)
	if err != nil {
		LogError(err)
		return
	}
	db := OpenDatabase()
	if db == nil {
		return
	}
	id := fmt.Sprintf("%x", md5.Sum([]byte(p.Title)))
	t := time.Now()
	stmt, err := db.Prepare("INSERT IGNORE INTO blog_posts(ID,TITLE,CONTENT,DATE,TIME) VALUES(?,?,?,?,?)")
	if err != nil {
		LogError(err)
		return
	}
	res, err := stmt.Exec(id, p.Title, p.Content,
		fmt.Sprintf("%04d/%02d/%02d", t.Year(), t.Month(), t.Day()),
		fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second()))
	//res, err := db.Exec("INSERT IGNORE INTO blog_posts (ID,TITLE,CONTENT,DATE,TIME) VALUES ('" + id + "','" + p.Title + "','" + p.Content + "','" + fmt.Sprintf("%04d/%02d/%02d", t.Year(), t.Month(), t.Day()) + "','" + fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second()) + "')")
	if err != nil {
		LogError(err)
		return
	}
	ro, _ := res.RowsAffected()
	LogIt(p, ro)
}

func render(w http.ResponseWriter, tmpl string, context Context) {
	context.Static = STATIC_URL
	tl := []string{TEMPLATES + "/base.tmpl", fmt.Sprintf(TEMPLATES+"/%s.tmpl", tmpl)}
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
	http.HandleFunc("/delete/", Delete)
	http.HandleFunc(STATIC_URL, Static)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
