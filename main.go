package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/secure"
	"net/http"
)

type Route struct {
	Name         string
	Method       string
	Pattern      string
	HandlerFunc  http.HandlerFunc
	Authenticate bool
}

type Routes []Route

var healthCheckRoute = Route{
	"HealthCheck",
	"GET",
	"/healthz",
	func(writer http.ResponseWriter, r *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("OK"))
	},
	false,
}

var port = "3333"

func main() {
	secureMiddleware := secure.New(secure.Options{IsDevelopment: true})
	r := mux.NewRouter()
	r.Use(secureMiddleware.Handler)
	r.Methods(healthCheckRoute.Method).
		Path(healthCheckRoute.Pattern).
		Name(healthCheckRoute.Name).
		Handler(healthCheckRoute.HandlerFunc)
	http.Handle("/", r)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		logrus.Errorln("An error occured starting HTTP listener at port " + port)
		logrus.Errorln("Error: " + err.Error())
	}
}
