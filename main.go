package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from CI/CD Pipeline in Go!")
    fmt.Println("testing pipline")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server running on port 80...")
    http.ListenAndServe(":80", nil)
}

