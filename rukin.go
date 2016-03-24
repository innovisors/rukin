package main

import (
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("C:\\Users\\Public")))
    http.ListenAndServe(":80", nil)
}
