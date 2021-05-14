# go-encode

Golang encoding utils.

```golang
import encode "github.com/wisdomfusion/go-encode"

...
e, _, _, err := encode.Find(resp.Body)
if err != nil {
    panic(err)
}

utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
...

```