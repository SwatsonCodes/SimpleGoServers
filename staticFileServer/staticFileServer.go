package main

import "net/http";

func main() {
  fileServer := http.FileServer(http.Dir("./static"))
  // in order to serve static files with a URL like "localhost:8000/static-files/content.txt" we need to strip the path prefix
  // http://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
  http.Handle("/static-files/", http.StripPrefix("/static-files/", fileServer))
  http.ListenAndServe(":8000", nil)
}
