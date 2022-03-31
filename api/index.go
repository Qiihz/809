package handler

import (
  "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  h := "<h1>Hello from Go!</h1>"
  fmt.Fprintf(w, h)
}
