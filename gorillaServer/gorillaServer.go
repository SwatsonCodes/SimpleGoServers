package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func printPathVar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "URL var is: %v\n", vars["someVar"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/path-var/{someVar}", printPathVar).Name("path")
	builtURL, _ := r.Get("path").URL("someVar", "foo")
	fmt.Printf("Sample built URL: %s", builtURL.String())
	http.ListenAndServe(":8000", r)
}
