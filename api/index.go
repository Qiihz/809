package handler

import (
  "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  r := "<h1>Hello from Go!</h1>"
  return r
}
