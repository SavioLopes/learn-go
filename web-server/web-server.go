package main

import (
  "fmt"
  "html"
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
  var err = http.ListenAndServe(":8080", nil)
  log.Fatal(err)
}