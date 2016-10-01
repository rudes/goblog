package main

import (
	"html/template"
	"time"
)

// Payload is the response
type Payload struct {
	ID, Title string
	Content   template.HTML
	Date      time.Time
}

// Post for post requests
type Post struct {
	Key            string
	Title, Content string
}

// Context for website
type Context struct {
	Payload []Payload
	Config  Config
}

// Config struct for TOML file
type Config struct {
	Title, Subtitle          string
	Author, AuthorEmail, Key string
	Tags, Description        string
}

// Message struct for response json
type Message struct {
	Message string
}

// ByDate allows sorting of Payload
type ByDate []Payload

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[j].Date.Before(a[i].Date) }
