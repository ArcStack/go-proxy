package main

import (
	"github.com/gorilla/mux"
	"github.com/lab/go-proxy/config"
	"github.com/lab/go-proxy/models"
	"github.com/lab/go-proxy/service"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/secure"
	"net/http"
)

var healthCheckRoute = models.Route{
	"HealthCheck",
	"GET",
	"/healthz",
	func(writer http.ResponseWriter, r *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("OK"))
	},
	false,
}

func main() {
	secureMiddleware := secure.New(secure.Options{IsDevelopment: true})
	r := mux.NewRouter()
	r.Use(secureMiddleware.Handler)
	r.Methods(healthCheckRoute.Method).
		Path(healthCheckRoute.Pattern).
		Name(healthCheckRoute.Name).
		Handler(healthCheckRoute.HandlerFunc)

	r.Methods(service.ReverseToFlaskRoute.Method).
		Path(service.ReverseToFlaskRoute.Pattern).
		Name(service.ReverseToFlaskRoute.Name).
		Handler(service.ReverseToFlaskRoute.HandlerFunc)

	http.Handle("/", r)
	err := http.ListenAndServe(":"+config.Port, nil)

	if err != nil {
		logrus.Errorln("An error occured starting HTTP listener at port " + config.Port)
		logrus.Errorln("Error: " + err.Error())
	}
}
