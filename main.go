package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    message := os.Getenv("MESSAGE")
    if message == "" {
        message = "Hello, World!"
    }
    fmt.Fprintf(w, message)
}

func main() {
    http.HandleFunc("/", handler)
    log.Println("Starting server on :9000")
    log.Fatal(http.ListenAndServe(":9000", nil))
}
