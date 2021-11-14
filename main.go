package main

import (
	"log"
	"net/http"
)

func main() {

	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		})
	*/

	http.Handle("/", http.FileServer(http.Dir("./static")))

	port := ":8080"

	log.Fatal(http.ListenAndServe(port, nil))

}
