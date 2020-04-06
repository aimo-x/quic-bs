package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/lucas-clemente/quic-go/http3"
)

func main() {
	client := &http.Client{
		Transport: &http3.RoundTripper{},
	}
	req, err := http.NewRequest("POST", "https://iuu.pub", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)
	if err != nil {
		// handle error
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
