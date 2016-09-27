package main

import (
	"errors"
	"html/template"
	"strings"
	"time"

	"github.com/gocql/gocql"
)

func openDB() (*gocql.Session, error) {
	cluster := gocql.NewCluster("db")
	cluster.Keyspace = "letters"
	cluster.Consistency = gocql.Quorum
	return cluster.CreateSession()
}

func post(db *gocql.Session, p Payload) error {
	if db == nil {
		return errors.New("Got nil cql session")
	}

	err := db.Query(`INSERT INTO letters ( id, title, content, date) VALUES (?, ?, ?, ?)`,
		p.ID, p.Title, p.Content, p.Date).Exec()
	if err != nil {
		return err
	}

	return nil
}

func getone(db *gocql.Session, postID string) ([]Payload, error) {
	if db == nil {
		return nil, errors.New("Got nil cql session")
	}

	var id, title, content string
	var date time.Time
	var p []Payload

	data := db.Query("SELECT id, title, content, date FROM letters WHERE id = ?", postID).Iter()
	defer data.Close()
	for data.Scan(&id, &title, &content, &date) {
		content = template.HTMLEscapeString(content)
		content = strings.Replace(content, "\n", "<br>", -1)
		p = append(p, Payload{
			ID:      id,
			Title:   title,
			Content: template.HTML(content),
			Date:    date,
		})
	}

	return p, nil
}

func getall(db *gocql.Session) ([]Payload, error) {
	if db == nil {
		return nil, errors.New("Got nil cql session")
	}

	var id, title, content string
	var date time.Time
	var p []Payload

	data := db.Query("SELECT id, title, content, date FROM letters").Iter()
	defer data.Close()
	for data.Scan(&id, &title, &content, &date) {
		content = template.HTMLEscapeString(content)
		content = strings.Replace(content, "\n", "<br>", -1)
		p = append(p, Payload{
			ID:      id,
			Title:   title,
			Content: template.HTML(content),
			Date:    date,
		})
	}

	return p, nil
}
