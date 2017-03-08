


package main

import (
  "fmt";
  "net/http";
  "net/url"
  "html/template"
)

type Name struct {
  Name string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, world!\n")
}

func printRequestMethod(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET":
      fmt.Fprintf(w, "You're using a GET request!\n")
    case "POST":
      fmt.Fprintf(w, "Recieved a POST!\n")
    default:
      fmt.Fprintf(w, "You sent a %s request\n", r.Method)
  }
}

func printRequestInfo(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Request method is: %s\n", r.Method)
  fmt.Fprintf(w, "URL is: %s\n", r.URL)
  fmt.Fprintf(w, "URL args are:\n")
  args, _ := url.ParseQuery(r.URL.RawQuery)
  for key := range args{
    fmt.Fprintf(w, "%s: %s\n", key, args[key])
  }
}

func templateName(w http.ResponseWriter, r *http.Request) {
  args, _ := url.ParseQuery(r.URL.RawQuery)
  name := &Name{args["name"][0]}
  t, _ := template.ParseFiles("./templates/howdy.html")
  t.Execute(w, name)
}

func main() {
  http.HandleFunc("/hello", sayHello)
  http.HandleFunc("/method", printRequestMethod)
  http.HandleFunc("/info", printRequestInfo)
  // no way to do variable parts to URL (with native library)!

  fileServer := http.FileServer(http.Dir("./static"))
  // in order to serve static files with a URL like "localhost:8000/static-files/content.txt" we need to strip the path prefix
  // http://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
  http.Handle("/static-files/", http.StripPrefix("/static-files/", fileServer))

  http.HandleFunc("/name", templateName)

  http.ListenAndServe(":8000", nil)
}
