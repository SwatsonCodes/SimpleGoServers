# SimpleGoServers

This is a handful of simple servers written in Go demonstrating useful patterns.
It covers most of the concepts in the [flask quickstart guide](http://flask.pocoo.org/docs/0.12/quickstart/), but (obviously)
translated to Go. All of the servers use built in packages, with the exception of GorillaServer which uses Gorilla Mux to 
demonstrate some features not present in the default 'net/http' library.

All of the servers run on `localhost:8000`. You can run each one by navigating to the server's directory and running `go run [server_name].go`. Each server contains a readme with a brief explanation and some sample requests.
