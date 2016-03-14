package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	data := json.NewDecoder(r.Body)
	var req Request
	err := data.Decode(&req)
	if err != nil {
		l.Log(err)
		return
	}
	res, err := json.Marshal(handleReq(&req))
	if err != nil {
		l.Log(err)
		return
	}
	fmt.Fprintf(w, "%s", string(res))
}

func main() {
	http.HandleFunc("/api/", handler)
	l.Log(http.ListenAndServe(":8080", nil))
}
