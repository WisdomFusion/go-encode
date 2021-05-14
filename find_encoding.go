package encode

import (
	"bufio"
	"io"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

func FindEncoding(r io.Reader) (e encoding.Encoding, name string, certain bool, err error) {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return
	}

	e, name, certain = charset.DetermineEncoding(bytes, "")
	return
}
