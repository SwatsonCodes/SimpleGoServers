package main

import (
  "fmt";
  "net/http";
  "net/url"
)

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

func main() {
  http.HandleFunc("/hello", sayHello)
  http.HandleFunc("/method", printRequestMethod)
  http.HandleFunc("/info", printRequestInfo)
  http.ListenAndServe(":8000", nil)
}
