package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
 	fmt.Fprintf(w, "I'm %s\n", hostname)
    })
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(os.Stdout, "Version 1")
 	fmt.Fprintf(w, "version 1")
    })



    log.Fatal(http.ListenAndServe(":" + port, nil))
}

