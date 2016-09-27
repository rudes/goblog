package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"sort"
	"time"
)

const (
	_staticURL    = "/static/"
	_templateRoot = "/go/src/github.com/rudes/otherletters.net/static/templates/"
	_staticRoot   = "/go/src/github.com/rudes/otherletters.net/static/"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/post/", postHandler)
	http.HandleFunc(_staticURL, staticHandler)
	http.ListenAndServe(":8080", nil)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	sf := r.URL.Path[len(_staticURL):]
	if len(sf) != 0 {
		f, err := http.Dir(_staticRoot).Open(sf)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, r, sf, time.Now(), content)
			return
		}
	}
	http.NotFound(w, r)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	conf, err := getConfig(_staticRoot + "config.toml")
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	var p Post
	b := make([]byte, 10000)
	r.Body.Read(b)
	json.Unmarshal(b, &p)
	if p.Key == conf.Key {
		p.Key = ""
		db, err := openDB()
		if err != nil {
			fmt.Fprintf(w, "%s", err)
		}
		err = post(db, p)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
		}
	}
	conf.Key = ""
	http.NotFound(w, r)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	post := r.URL.Path[len("/view/"):]
	db, err := openDB()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	p, err := getone(db, post)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	sort.Sort(ByDate(p))
	sort.Sort(ByTime(p))

	render(w, p)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	db, err := openDB()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	p, err := getall(db)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	sort.Sort(ByDate(p))
	sort.Sort(ByTime(p))

	render(w, p)
}

func render(w http.ResponseWriter, payload []Payload) {
	conf, err := getConfig(_staticRoot + "config.toml")
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	context := Context{
		Config:  conf,
		Payload: payload,
	}

	tl := []string{_templateRoot + "base.tmpl",
		_templateRoot + "header.tmpl",
		_templateRoot + "index.tmpl",
	}

	t, err := template.ParseFiles(tl...)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	err = t.Execute(w, context)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
}
