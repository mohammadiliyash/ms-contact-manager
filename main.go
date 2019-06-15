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
	//fmt.Printf("config: %v\n", cfg)

	n := negroni.New()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.Use(negroni.NewRecovery())
	//nlog.LogRequestStart = false
	//n.Use(nlog.New("MSL"))
	//n.Use(c)
	n.UseHandler(routes.New(handler{}))
	n.Run(":4000")

}

type handler struct{}

// ServeHTTP is the implementation of the http.Handler interface. It serves. diagnostics page requests.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// if r == nil {
	// 	// not much we can do here
	// 	logger.Printf("Diagnostics: Nil request sent to ServeHTTP")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	//url := r.URL
	// if url == nil {
	// 	if logger != nil {
	// 		logger.Printf("Diagnostics: Nil URL sent to ServeHTTP")
	// 	}
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// q := url.Query().Get(runTestsParameterName)

	code := http.StatusOK
	//runTests := q != "" && q != "false"

	// result := getCurrentSystemInfo(runTests)
	// if runTests && result.FailedTests != nil && *result.FailedTests > 0 {
	// 	code = http.StatusInternalServerError
	// }

	//resultJSON, err := json.MarshalIndent(result, "", "  ")
	// if err != nil {
	// 	if logger != nil {
	// 		logger.Printf("Diagnostics: Error serializing result set:\n%v", err.Error())
	// 	}
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	//_, _ = w.Write(resultJSON)
	// if err != nil && logger != nil {
	// 	logger.Printf("Diagnostics: Error writing result set:\n%v", err.Error())
	// }
}
