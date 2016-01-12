package main

import (
	"log"
	"os"
)

func LogIt(t Post, s int64) {
	f, err := os.OpenFile("/var/log/OtherLetterApi.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	if s == 0 {
		log.Println(t.Title + " Already Exists...")
	} else {
		log.Println(t.Title + " Added")
	}
}
