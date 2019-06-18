package main

import (
	"net/http"

	"github.com/miliyash/ms-contact-manager/app/routes"
	cfg "github.com/miliyash/ms-contact-manager/config"
	log "github.com/miliyash/ms-contact-manager/logging"
	gzip "github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"
)

const LISTEN_PORT = "4000"

func main() {

	log.Info("Service initialized")

	log.New().WithFields(map[string]interface{}{
		"e":    "webserver_start",
		"port": LISTEN_PORT,
	}).Info(".")

	cfg.Initialize()

	n := negroni.New()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.Use(negroni.NewRecovery())

	n.UseHandler(routes.New(handler{}))
	n.Run(":" + LISTEN_PORT)

}

type handler struct{}

// ServeHTTP is the implementation of the http.Handler interface. It serves. diagnostics page requests.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	code := http.StatusOK
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

}
