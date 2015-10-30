# go-sd
Go library for systemd

## Socket Activation

```go
import "github.com/victoranator/go-sd"

func main() {
    ls, _ := sd.Listeners()
    if len(ls) > 0 {
        for i, l := range ls {
            fmt.Printf("%d; Name: %s; Addr: %s\n", i, l.Name(), l.Addr().String())
            http.Serve(l, nil)
            defer l.Close()
        }
    }
}
```
