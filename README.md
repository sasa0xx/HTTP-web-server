# HTTP-web-server

## An EXAMPLE project that shows how to build a simple HTTP web server using go

This project is a simple example of building an HTTP server using the go *net/http* header only and no depenencies
This project is not meant to be used in real-world scenarios, it is only for example and educational purposes
This project also uses a dummy map in memory instead of an actual dataset like MySQL

The project demonstrates how to:
- Listen to an HTTP port like 8080
- Create GET, POST and DELETE endpoints for the port
- Manage a memory map and make it thread-safe
- Save and retrive data in the form of JSON

## Features

- Simple HTTP server listening on port :8080
- GET, POST and DELETE endpoints
- A simple thread-safe memory map to act as the dataset

## Installation & Usage

To install the package for use in your project:

```bash
go get github.com/sasa0xx/HTTP-web-server
```

Or add it to your module manually:

```bash
go mod edit -require=github.com/sasa0xx/HTTP-web-server
```

To test the project you can use Postman or the simply the curl command, here are some test cases and outputs:

```bash
curl -i http://localhost:8080/
```
Output:
```
200 OK
Content-Type: text/plain
Body: Hello, World! :D
```

```bash
curl -i -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{}'
```
Output:
```
400 Bad Request
Body: User name is required
```

```bash
curl -i -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice"}'
```
Output:
```
204 No content
```

```bash
curl -i http://localhost:8080/users/0
```
Output:
```
200 OK
Content-Type: application/json
Body: {"name":"Alice"}
```

```bash
curl -i http://localhost:8080/users/999
```
Output:
```
400 Bad Request
User not found.
```

```bash
curl -i -X DELETE http://localhost:8080/users/0
```
Output:
```
204 No Content
```


```bash
curl -i http://localhost:8080/users/0
```
Output:
```
400 Bad Request
User not found.
```
