package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func recieveFile(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var Buffer bytes.Buffer
	io.Copy(&Buffer, file)
	fmt.Fprintf(w, "Contents of file %s: %v", header.Filename, string(Buffer.Bytes()))
}

func main() {
	http.HandleFunc("/upload", recieveFile)
	http.ListenAndServe(":8000", nil)
}
