package main

import (
	"context"
	"net/http"

	"github.com/lucas-clemente/quic-go/http3"
)

var ctx = context.Background()

func main() {
	http.Handle("/", http.FileServer(http.Dir("./www")))
	err := http3.ListenAndServeQUIC(":443", "./../CSR/2822297_iuu.pub.pem", "./../CSR/2822297_iuu.pub.key", nil)
	if err != nil {
		panic(err)
	}
}
