package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Microsecond}
	resp, err := c.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// ioutil deprecated: Go 1.16
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
