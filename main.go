package main

import (
        "fmt"
        "log"
        "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
        title := "Jenkins X golang http example"

        from := ""
        if r.URL != nil {
                from = r.URL.String()
        }
        if from != "/favicon.ico" {
                log.Printf("title: %s\n", title)
        }

        fmt.Fprintf(w, "<html><body><img style='width:100%%' src='/static/download-banner.png'></body></html>")
}

func main() {
        http.HandleFunc("/", handler)
        fs := http.FileServer(http.Dir("./static"))
        http.Handle("/static/", http.StripPrefix("/static", fs))
        http.ListenAndServe(":8080", nil)
}

