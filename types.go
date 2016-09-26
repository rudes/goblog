package main

import "html/template"

// Payload is the response
type Payload struct {
	ID, Title  string
	Content    template.HTML
	Date, Time string
}

// Context for website
type Context struct {
	Payload []Payload
	Config  Config
}

// Config struct for TOML file
type Config struct {
	Title, Subtitle, Author, AuthorEmail string
}

// ByDate allows sorting of Payload
type ByDate []Payload

// ByTime allows sorting of Payload
type ByTime []Payload

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Date < a[j].Date }

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return a[i].Time < a[j].Time }
