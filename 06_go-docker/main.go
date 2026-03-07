package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	println("Server listening on port 8081")
	http.ListenAndServe(":8081", nil)
}
