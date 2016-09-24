package main

import (
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
	http.HandleFunc("/", handler)
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

func handler(w http.ResponseWriter, r *http.Request) {
	p, err := getall()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	sort.Sort(ByDate(p))

	render(w, p)
}

func render(w http.ResponseWriter, payload []Payload) {
	tl := []string{_templateRoot + "base.tmpl",
		_templateRoot + "header.tmpl",
		_templateRoot + "index.tmpl",
	}

	t, err := template.ParseFiles(tl...)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	err = t.Execute(w, payload)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
}
