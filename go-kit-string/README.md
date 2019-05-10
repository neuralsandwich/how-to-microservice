# Go Kit Simple String

Blog post: 

This is an example program from a blog series about building microservices in
Go.

## What do I need to have a play?

* You will need to [install Go](https://golang.org/dl/)
* You will need either `wget` or `curl`

## How do I build it?

```
go build
```

## How do I run it?

```
$ ./simple-string &

$ curl -d '{"s": “Hello, World!"}' http://localhost:8080/count
> {"c":13}

$ curl -d '{"s": "Hello, World!"}' http://localhost:8080/uppercase
{"s":"HELLO, WORLD!"}
```
