package main

import (
	"database/sql"

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

func GetAllLetters() []Post {
	var p []Post
	db := OpenDatabase()
	if db == nil {
		return nil
	}
	rows, err := db.Query("SELECT ID,TITLE,CONTENT,DATE,TIME FROM blog_posts ORDER BY DATE DESC, TIME DESC")
	if err != nil {
		LogError(err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var po Post
		err := rows.Scan(&po.ID, &po.Title, &po.Content, &po.Date, &po.Time)
		if err != nil {
			LogError(err)
			return nil
		}
		p = append(p, po)
	}
	return p
}
