package main

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func handleReq(req *Request) Response {
	var res Response
	var err error
	var p Post
	switch req.Action {
	case "get":
		p, err = getLetter(req.Post.ID)
		res.Posts = []Post{p}
		res.Message = fmt.Sprintf("%v", err)
	case "all":
		res.Posts, err = getAllLetters()
		res.Message = fmt.Sprintf("%v", err)
	case "delete":
		err = deleteLetter(req.Post.ID)
		res.Message = fmt.Sprintf("%v", err)
	case "updated":
		err = updateLetter(req.Post)
		res.Message = fmt.Sprintf("%v", err)
	case "new":
		err = newLetter(req.Post)
		res.Message = fmt.Sprintf("%v", err)
	default:
		res.Message = "Unknown Command: " + req.Action
	}
	if err != nil {
		l.Log(err)
		res.Message = fmt.Sprintf("Error Detected: %s", res.Message)
	}
	return res
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "CHANGEME")
	if err != nil {
		l.Log(err)
		return nil
	}
	return db
}

func getLetter(postID string) (Post, error) {
	var p Post
	db := openDB()
	if db == nil {
		return p, errors.New("Could Not Open Database")
	}
	l.Log("Getting post: ", postID)
	err := db.QueryRow("SELECT ID,TITLE,CONTENT,DATE,TIME FROM posts WHERE ID="+
		postID).Scan(&p.ID,
		&p.Title,
		&p.Content,
		&p.Date,
		&p.Time)
	if err != nil {
		return p, err
	}
	return p, nil
}

func getAllLetters() ([]Post, error) {
	var p []Post
	db := openDB()
	if db == nil {
		return nil, errors.New("Could Not Open Database")
	}
	l.Log("Getting all posts.")
	rows, err := db.Query("SELECT ID,TITLE,CONTENT,DATE,TIME FROM posts ORDER BY DATE DESC, TIME DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var po Post
		err := rows.Scan(&po.ID, &po.Time, &po.Content, &po.Date, &po.Time)
		if err != nil {
			return nil, err
		}
		p = append(p, po)
	}
	return p, nil
}

func updateLetter(p Post) error {
	db := openDB()
	if db == nil {
		return errors.New("Could Not Open Database")
	}
	defer db.Close()
	l.Log("Updating post: ", p.ID)
	stmt, err := db.Prepare("UPDATE posts SET CONTENT=? WHERE ID=" + p.ID)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(p.Content)
	if err != nil {
		return err
	}
	ro, _ := res.RowsAffected()
	if ro == 0 {
		l.Log(p.Title, " Not Updated...")
	} else {
		l.Log(p.Title, " Updated")
	}
	return nil
}

func deleteLetter(postID string) error {
	db := openDB()
	if db == nil {
		return errors.New("Could Not Open Database")
	}
	defer db.Close()
	l.Log("Deleting post: ", postID)
	res, err := db.Exec("DROP FROM posts WHERE ID=" + postID)
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows != 0 {
		l.Log("Deleted Post: ", postID)
	} else {
		l.Log("Post Not Deleted: ", postID)
	}
	return nil
}

func newLetter(p Post) error {
	db := openDB()
	if db == nil {
		return errors.New("Could Not Open Database")
	}
	defer db.Close()
	l.Log("Creating New Post: ", p.ID)
	id := fmt.Sprintf("%x", md5.Sum([]byte(p.Title)))
	t := time.Now()
	stmt, err := db.Prepare("INSERT IGNORE INTO posts(ID,TITLE,CONTENT,DATE,TIME) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id, p.Title, p.Content,
		fmt.Sprintf("%04d/%02d/%02d", t.Year(), t.Month(), t.Day()),
		fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second()))
	if err != nil {
		return err
	}
	ro, _ := res.RowsAffected()
	if ro == 0 {
		l.Log(p.Title, " Already Exists...")
	} else {
		l.Log(p.Title, " Added")
	}
	return nil
}
