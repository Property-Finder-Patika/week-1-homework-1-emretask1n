
# Exercise 3.4

Following the approach of the `Lissajous` example in Section 1.7, construct a web server that computes surfaces and writes SVG data to the client. The server must set the `Content-Type` header like this:

```go
w.Header().Set("Content-Type", "image/svg+xml")
```