package main

import (
	"github.com/unrolled/secure"
	"net/http"
)

var myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello world"))
	if err != nil {
		return
	}
})

func main() {
	secureMiddleware := secure.New(secure.Options{IsDevelopment: true})
	app := secureMiddleware.Handler(myHandler)
	err := http.ListenAndServe("127.0.0.1:3333", app)
	if err != nil {
		return
	}
}
