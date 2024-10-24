package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format(time.RFC1123)
    fmt.Fprintf(w, "Current Date and Time: %s", currentTime)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}