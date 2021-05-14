# go-encode

Golang encoding utils.

```golang
import encode "github.com/wisdomfusion/go-encode"

...
e, _, _, err := encode.FindEncoding(resp.Body)
if err != nil {
    panic(err)
}

utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
...

```