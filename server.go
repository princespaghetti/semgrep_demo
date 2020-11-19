package main

import (
    "fmt"
    "net/http"
)

func greeting(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}


func main() {

    http.HandleFunc("/greeting", greeting)

    http.ListenAndServe(":80", nil)
}