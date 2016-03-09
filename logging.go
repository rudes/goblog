package main

import (
	"log"
	"os"
)

func openFile() *os.File {
	f, _ := os.OpenFile("/var/log/OtherLetterApi.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return f
}

func LogAnything(t ...interface{}) {
	f := openFile()
	defer f.Close()
	log.SetOutput(f)
	log.Println(t...)
}

func LogIt(t Post, s int64) {
	f := openFile()
	defer f.Close()
	log.SetOutput(f)
	if s == 0 {
		log.Println(t.Title + " Already Exists...")
	} else {
		log.Println(t.Title + " Added")
	}
}

func LogError(err error) {
	f := openFile()
	defer f.Close()
	log.SetOutput(f)

	log.Println(err)
}
