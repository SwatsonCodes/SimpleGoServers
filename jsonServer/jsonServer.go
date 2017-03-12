package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type user struct {
	Name        string
	Description string
}

type userLookup struct {
	Users map[string]user
}

func (lookup *userLookup) handleUser(w http.ResponseWriter, r *http.Request) {
	args, _ := url.ParseQuery(r.URL.RawQuery)
	name, namePresent := args["name"]
	switch r.Method {
	case "GET":
		if namePresent {
			userName := name[0]
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(lookup.Users[userName])
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error: must supply user name")
		}
	case "POST":
		{
			decoder := json.NewDecoder(r.Body)
			var u user
			err := decoder.Decode(&u)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Error: unable to decode json payload")
				return
			}
			if u.Name == "" || u.Description == "" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Error: must specify user name and description")
				return
			}
			lookup.Users[u.Name] = u
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "OK")
		}
	default:
		fmt.Fprintf(w, "Unsupported request type: %s\n", r.Method)
	}
}

func main() {
	lookup := userLookup{make(map[string]user)}
	http.HandleFunc("/user", lookup.handleUser)
	http.ListenAndServe(":8000", nil)
}
