package webserver

import (
	"net/http"

	db "github.com/miliyash/ms-contact-manager/app/database"
	"github.com/miliyash/ms-contact-manager/app/routes"
	log "github.com/miliyash/ms-contact-manager/logging"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"
)

func NewServer() *negroni.Negroni {
	log.Info("Service initialized")

	n := negroni.New()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.Use(negroni.NewRecovery())

	n.UseHandler(routes.New(handler{}))

	db.Initialize()

	return n
}

type handler struct{}

// ServeHTTP is the implementation of the http.Handler interface. It serves. diagnostics page requests.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	code := http.StatusOK
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

}
