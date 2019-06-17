package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/miliyash/ms-contact-manager/app/routes"
	cfg "github.com/miliyash/ms-contact-manager/config"
	log "github.com/miliyash/ms-contact-manager/logging"
	gzip "github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"
)

func main() {

	cfg.Initialize()
	log.AppName = cfg.AppName
	fmt.Println(cfg.AppName)
	err := errors.New("Error")
	log.WithError(err)
	log.Error(err, nil)
	n := negroni.New()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.Use(negroni.NewRecovery())

	n.UseHandler(routes.New(handler{}))
	n.Run(":4000")

}

type handler struct{}

// ServeHTTP is the implementation of the http.Handler interface. It serves. diagnostics page requests.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	code := http.StatusOK

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

}
