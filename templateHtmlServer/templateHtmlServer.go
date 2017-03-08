package main

import (
  "net/http";
  "net/url"
  "html/template"
)

type Name struct {
  Name string
}

func templateName(w http.ResponseWriter, r *http.Request) {
  args, _ := url.ParseQuery(r.URL.RawQuery)
  name := &Name{args["name"][0]}
  t, _ := template.ParseFiles("./templates/howdy.html")
  t.Execute(w, name)
}

func main() {
  http.HandleFunc("/name", templateName)
  http.ListenAndServe(":8000", nil)
}
