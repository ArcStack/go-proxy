package service

import "github.com/lab/go-proxy/models"

var ReverseToFlaskRoute = models.Route{
	Name:         "ReverseToGitHub",   // Name
	Method:       "GET",               // HTTP method
	Pattern:      "/reverse-to-flask", // Route pattern
	HandlerFunc:  ReverseToFlask,
	Authenticate: true,
}
