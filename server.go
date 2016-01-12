package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Structure for json data
type Post struct {
	Title, Content string
}

// Handle incoming post data
// Submit data into sql
func HandleLetters(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var p Post
	err := d.Decode(&p)
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("mysql", "blog:a40echo14b19@/BLOG")
	if err != nil {
		log.Fatal(err)
	}
	id := fmt.Sprintf("%x", md5.Sum([]byte(p.Title)))
	t := time.Now()
	res, err := db.Exec("INSERT IGNORE INTO blog_posts (ID,TITLE,CONTENT,DATE,TIME) VALUES ('" + id + "','" + p.Title + "','" + p.Content + "','" + fmt.Sprintf("%04d/%02d/%02d", t.Year(), t.Month(), t.Day()) + "','" + fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second()) + "')")
	if err != nil {
		log.Fatal(err)
	}
	ro, _ := res.RowsAffected()
	LogIt(p, ro)
}

func main() {
	http.HandleFunc("/api", HandleLetters)
	log.Fatal(http.ListenAndServe(":80", nil))
}
