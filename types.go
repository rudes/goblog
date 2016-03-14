package main

import "github.com/rudes/crylog"

var l crylog.CryLog

func init() {
	l.File = "/var/www/OtherLetterAPI.log"
}

type Post struct {
	ID, Title, Content, Date, Time string
}

type Request struct {
	PostID, Action string
}

type Response struct {
	Posts   []Post
	Message string
}
