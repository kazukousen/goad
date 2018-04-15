package main

import (
	"net/http"

	_ "github.com/kazukousen/goad/api/handler"
)

func campaignHandle() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// do something ...
		w.Write([]byte("test"))
	}
}

func main() {
	http.Handle("/test/", http.HandlerFunc(campaignHandle()))
	http.ListenAndServe(":8080", nil)
}
