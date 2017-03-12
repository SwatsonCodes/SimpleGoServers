This server uses Gorilla Mux to demonstrate a couple useful features not contained in the standared `net/http` library.
Namely, it shows 1) using code to automatically generate a URL and 2) reading variables in a URL path.

When you start the server (`go run gorillaServer.go`) it will print out a sample path it automatically built.
This will make more sense if you look at the code. Then try running
`curl -X GET 'localhost:8000/path-var/some_variable'`
