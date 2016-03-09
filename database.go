package main

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDatabase() *sql.DB {
	db, err := sql.Open("mysql", "blog:a40echo14b19@/BLOG")
	if err != nil {
		LogError(err)
		return nil
	}
	return db
}

func DeleteThisLetter(postID string) error {
	db := OpenDatabase()
	defer db.Close()
	LogAnything("Deleting post: ", postID)
	res, err := db.Exec("DROP FROM blog_posts WHERE ID=" + postID)
	if rows, _ := res.RowsAffected(); rows != 0 {
		LogAnything("Deleted Post: ", postID)
	}
	return err
}

func GetThisLetter(postID string) (*Post, error) {
	var p Post
	db := OpenDatabase()
	LogAnything("Retrieving post: ", postID)
	err := db.QueryRow("SELECT ID,TITLE,CONTENT,DATE,TIME FROM blog_posts WHERE ID="+
		postID).Scan(&p.ID,
		&p.Title,
		&p.Content,
		&p.Date,
		&p.Time)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func GetAllLetters() ([]Post, error) {
	var p []Post
	db := OpenDatabase()
	if db == nil {
		return nil, errors.New("Could Not Open Database")
	}
	LogAnything("Retrieving all posts")
	rows, err := db.Query("SELECT ID,TITLE,CONTENT,DATE,TIME FROM blog_posts ORDER BY DATE DESC, TIME DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var po Post
		err := rows.Scan(&po.ID, &po.Title, &po.Content, &po.Date, &po.Time)
		if err != nil {
			return nil, err
		}
		p = append(p, po)
	}
	return p, nil
}

func UpdateThisLetter(p Post) error {
	db := OpenDatabase()
	if db == nil {
		return errors.New("Could Not Open Database")
	}
	defer db.Close()
	LogAnything("Updating post: ", p.ID)
	stmt, err := db.Prepare("UPDATE blog_posts SET CONTENT=? WHERE ID=" + p.ID)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(p.Content)
	if err != nil {
		return err
	}
	ro, _ := res.RowsAffected()
	LogIt(p, ro)
	return nil
}

func HandleThisLetter(p Post) {
	db := OpenDatabase()
	if db == nil {
		return
	}
	LogAnything("Createing New Post: ", p.ID)
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
