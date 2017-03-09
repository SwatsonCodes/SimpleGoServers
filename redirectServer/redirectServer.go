package main

import "net/http"

func redirectEndpoint(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://www.purple.com", 301)
}

func main() {
	http.HandleFunc("/redirect", redirectEndpoint)
	http.ListenAndServe(":8000", nil)
}
