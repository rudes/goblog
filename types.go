package main

// Payload is the response
type Payload struct {
	ID, Title, Content, Date, Time string
}

// ByDate allows sworting of Payload
type ByDate []Payload

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Date < a[j].Date }
