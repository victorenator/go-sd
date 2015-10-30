= go-sd =

== Socket Activation ==

```golang
import "github.com/victoranator/go-sd"

func main() {
    ls, _ := sd.Listeners()
    if len(ls) > 0 {
        for i, l := range ls {
            fmt.Printf("%d Name: %+v; Addr: %v\n", i, l.Name(), l.Addr().String())
            http.Serve(l, nil)
            defer l.Close()
        }
    }
}
```
