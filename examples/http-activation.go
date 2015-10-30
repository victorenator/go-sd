package main

import "github.com/victorenator/go-sd"
import "fmt"
import "net/http"
import "io"

func RootResource(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello\n");
}

func main() {
    http.HandleFunc("/", RootResource)

    ls, _ := sd.Listeners()
    if len(ls) > 0 {
        for i, l := range ls {
            fmt.Printf("%d; Name: %s; Addr: %s\n", i, l.Name(), l.Addr().String())
            http.Serve(l, nil)
            defer l.Close()
        }

    } else {
        http.ListenAndServe(":8088", nil)
    }
}
