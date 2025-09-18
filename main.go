package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "testing pipline!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server running on port 80...")
    http.ListenAndServe(":80", nil)
}

