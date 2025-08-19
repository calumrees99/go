package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	// Serve static files (PDF goes here)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// CV page handler
	http.HandleFunc("/cv", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "cv.html")
	})

	// Start server
	fmt.Println("Server is listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
