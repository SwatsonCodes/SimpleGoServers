package main

import (
	"html/template"
	"net/http"
	"net/url"
)

type name struct {
	Name string
}

func templateName(w http.ResponseWriter, r *http.Request) {
	args, _ := url.ParseQuery(r.URL.RawQuery)
	n := &name{args["name"][0]}
	t, _ := template.ParseFiles("./templates/howdy.html")
	t.Execute(w, n)
}

func main() {
	http.HandleFunc("/name", templateName)
	http.ListenAndServe(":8000", nil)
}
