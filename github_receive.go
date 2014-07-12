package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(req.Form)
}

func main() {
	http.HandleFunc("/pullreq_notifier", handler)
	http.ListenAndServe(":8080", nil)
}
