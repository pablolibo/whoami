package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
)

func main() {
    version := "1.1"
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    fmt.Fprintf(os.Stdout, "start :%s\n", version)

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)

    hostname, _ := os.Hostname()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
 	fmt.Fprintf(w, "I'm %s\n", hostname)
    })

    http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(os.Stdout, "%s\n", version)
 	fmt.Fprintf(w, "%s\n", version)
    })

    log.Fatal(http.ListenAndServe(":" + port, nil))
}
