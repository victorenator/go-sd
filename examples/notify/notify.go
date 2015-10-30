package main

import "github.com/victorenator/go-sd"
import "net/http"
import "io"
import "time"

func RootResource(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello\n");
}

func main() {
    http.HandleFunc("/", RootResource)

    sd.Notify("STATUS=Starting ...")
    time.Sleep(10 * time.Second)
    sd.Notify("READY=1\nSTATUS=Ready to process ...")
    http.ListenAndServe(":8088", nil)
}
