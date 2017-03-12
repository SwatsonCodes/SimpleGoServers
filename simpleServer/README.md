This is a very basic server. It demonstrates receiving different types of requests and reading URL arguments.

Sample queries:
```
curl -X GET 'localhost:8000/hello'
curl -X GET 'localhost:8000/method'
curl -X POST 'localhost:8000/method'
curl -X GET 'localhost:8000/info?key1=val1&key2=val2&key2=val3'
```
