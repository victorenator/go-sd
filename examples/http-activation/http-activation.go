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

    fds := sd.Listeners()
    for i, fd := range fds {
        fmt.Printf("%d; Name: %s\n", i, fd.Name())
        l, _ := fd.Listener()
        if i < len(fds) - 1 {
            go func() {
                http.Serve(l, nil)
            }()
        } else {
            http.Serve(l, nil)
        }
    }
}
