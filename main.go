package main

import (
	"net/http"

	cfg "github.com/miliyash/ms-contact-manager/config"
	log "github.com/miliyash/ms-contact-manager/logging"
	web "github.com/miliyash/ms-contact-manager/web_server"
)

const LISTEN_PORT = "4000"

func main() {

	cfg.Initialize()
	server := web.NewServer()
	log.Info("Web server started")
	server.Run(":" + LISTEN_PORT)

}

type handler struct{}

// ServeHTTP is the implementation of the http.Handler interface. It serves. diagnostics page requests.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	code := http.StatusOK
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

}
