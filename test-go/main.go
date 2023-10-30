package main

import "fmt"
import "net/http"

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", r.URL.Path)
  })

  http.ListenAndServe(":8080", nil)
}
