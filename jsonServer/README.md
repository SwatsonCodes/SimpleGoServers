This server covers a few concepts, including:
* Reading a POST body
* Reading/writing JSON
* Sending a header status code

This server keeps track of "users" and their associated "descriptions." You can POST users and descriptions and later GET
users by their name.
Sample queries:
```
curl -X POST -H "Content-Type: application/json" -d '{"name":"Alice", "description": "engineer"}' 'localhost:8000/user'
curl -X GET 'localhost:8000/user?name=Alice'
curl -X POST -H "Content-Type: application/json" -d '{"name":"Bob"}' 'localhost:8000/user' -i // Bad request: must supply description
```
