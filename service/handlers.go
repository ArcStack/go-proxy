package service

import (
	"github.com/lab/go-proxy/reverseproxy"
	"net/http"
	"net/url"
)

func Reverse(w http.ResponseWriter, r *http.Request) {
	path, err := url.Parse("https://github.com")
	if err != nil {
		panic(err)
		return
	}
	proxy := reverseproxy.NewReverseProxy(path)
	proxy.ServeHTTP(w, r)
}
