package main

import "github.com/victorenator/go-sd"
import "io"
import "net/http"
import "os"
import "os/signal"
import "syscall"
import "time"

func WaitForTerminate() {
    signalCh := make(chan os.Signal, 1)
    signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
    <- signalCh
}

func main() {
    sd.Notify("STATUS=Starting ...")
    
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "hello\n");
    })

    time.Sleep(2 * time.Second)
    
    sd.Notify("READY=1\nSTATUS=Ready to process ...")
    go func() {
        http.ListenAndServe(":8088", nil)
    }()
    
    WaitForTerminate()
    
    sd.Notify("STOPPING=1\nSTATUS=Stopping ...")
    time.Sleep(2 * time.Second)
}
