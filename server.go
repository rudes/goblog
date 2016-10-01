package main

import (
	"crypto/md5"
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
	if r.Method == "GET" {
		render(w, nil, "post")
	} else {
		var postd Post
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&postd)
		id := md5.Sum([]byte(postd.Title + postd.Content))
		p := Payload{
			ID:      fmt.Sprintf("%x", id)[1:10],
			Title:   postd.Title,
			Content: template.HTML(postd.Content),
			Date:    time.Now().Local(),
		}
		fmt.Println(postd.Key, conf.Key)
		if postd.Key == conf.Key {
			fmt.Println(p.ID, p.Content)
			postd.Key = ""
			conf.Key = ""
			db, err := openDB()
			if err != nil {
				jsonRes(w, Message{
					Message: fmt.Sprintf("Error Opening DB: %s",
						err),
				})
				return
			}
			err = post(db, p)
			if err != nil {
				jsonRes(w, Message{
					Message: fmt.Sprintf("Error Posting: %s",
						err),
				})
				return
			}
			jsonRes(w, Message{
				Message: "Posted Successfully",
			})
			return
		}
		http.NotFound(w, r)
	}
}

func jsonRes(w http.ResponseWriter, message Message) error {
	res, err := json.Marshal(message)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return nil
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

	render(w, p, "index")
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

	render(w, p, "index")
}

func render(w http.ResponseWriter, payload []Payload, page string) {
	conf, err := getConfig(_staticRoot + "config.toml")
	conf.Key = ""
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	context := Context{
		Config: conf,
	}
	tl := []string{
		_templateRoot + "base.tmpl",
		_templateRoot + "header.tmpl",
		_templateRoot + page + ".tmpl",
	}
	context.Payload = payload

	tFuncs := template.FuncMap{
		"fmtDate": func(t time.Time) string {
			return t.Format("01/02/2006")
		},
		"recent": func() string {
			return fmt.Sprintf("%s",
				context.Payload[0].Date)
		},
	}
	t, err := template.New("base.tmpl").Funcs(tFuncs).ParseFiles(tl...)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	err = t.Execute(w, context)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	return
}
