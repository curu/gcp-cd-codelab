package main

import (
    "io"
    "log"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello, world!")
}

func busyloop(){
    for {
            log.Printf("hello")
    }
}

func main(){
    http.HandleFunc("/", hello)
    go busyloop()
    log.Fatal(http.ListenAndServe(":80", nil))
}
