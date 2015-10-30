# go-sd
Go library for systemd

## Socket Activation

```go
import "github.com/victorenator/go-sd"

func main() {
    fds, _ := sd.Listeners()
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
```

## Type=notify

```go
import "github.com/victorenator/go-sd"

func main() {
    sd.Notify("STATUS=Starting ...\n")
    time.Sleep(10 * time.Second)
    sd.Notify("READY=1\nSTATUS=Ready to process ...\n")
    http.ListenAndServe(":8088", nil)
}
```
