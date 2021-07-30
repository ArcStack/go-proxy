package service

import "github.com/lab/go-proxy/models"

var ReverseRoute = models.Route{
	Name:         "ReverseToGitHub", // Name
	Method:       "GET",             // HTTP method
	Pattern:      "/wjjmjh",         // Route pattern
	HandlerFunc:  Reverse,
	Authenticate: true,
}
