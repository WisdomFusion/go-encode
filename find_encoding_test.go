package encode

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFind(t *testing.T) {
	want := "gbk"

	resp, err := http.Get("https://www.qq.com")
	if err != nil {
		t.Errorf("Exception ocurred: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("HTTP request failed: status code", resp.StatusCode)
		return
	}

	_, name, _, err := FindEncoding(resp.Body)
	if err != nil {
		t.Errorf("Exception ocurred: %s", err.Error())
	}

	if got := name; got != want {
		t.Errorf("Got %s, but expect %s", got, want)
	}
}
