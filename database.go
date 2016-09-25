package main

import (
	"html/template"
	"strings"

	"github.com/gocql/gocql"
)

func openDB() (*gocql.Session, error) {
	cluster := gocql.NewCluster("db")
	cluster.Keyspace = "letters"
	cluster.Consistency = gocql.Quorum
	return cluster.CreateSession()
}

func getall() ([]Payload, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}

	var id, title, content, date, time string
	var p []Payload

	data := db.Query("SELECT id, title, content, date, time FROM letters").Iter()
	for data.Scan(&id, &title, &content, &date, &time) {
		content = template.HTMLEscapeString(content)
		content = strings.Replace(content, "\n", "<br>", -1)
		p = append(p, Payload{
			ID:      id,
			Title:   title,
			Content: template.HTML(content),
			Date:    date,
			Time:    time,
		})
	}

	if err := data.Close(); err != nil {
		return nil, err
	}

	return p, nil
}
