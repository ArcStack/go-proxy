package service

import (
	"github.com/lab/go-proxy/reverseproxy"
	"net/http"
	"net/url"
)

func ReverseToFlask(w http.ResponseWriter, r *http.Request) {
	path, err := url.Parse("http://0.0.0.0:5000/api")
	if err != nil {
		panic(err)
		return
	}
	proxy := reverseproxy.NewReverseProxy(path)
	proxy.ServeHTTP(w, r)
}
