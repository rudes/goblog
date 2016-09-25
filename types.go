package main

// Payload is the response
type Payload struct {
	ID, Title, Content, Date, Time string
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

// ByDate allows sworting of Payload
type ByDate []Payload

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Date < a[j].Date }
