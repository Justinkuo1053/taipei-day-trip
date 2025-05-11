package main

import (
    "log"
    "net/http"
)

func main() {
    log.Println("Starting server...")
    http.ListenAndServe(":8080", nil)
}