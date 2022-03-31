package handler

import (
  "fmt"
  "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  return "<h1>Hello from Go!</h1>"
}
