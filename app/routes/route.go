package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	c "github.com/miliyash/ms-contact-manager/app/controllers"
	cfg "github.com/miliyash/ms-contact-manager/config"
)

// New initializes routes for the app
func New(diagnosticsHandler http.Handler) *mux.Router {

	pathPrefixAppName := pathHeartbeat + cfg.AppName

	// Create main router
	mainRouter := mux.NewRouter().StrictSlash(true)

	//requestHandler := new(utils.RequestHandler)
	//accessService := uaa.NewAccessService(requestHandler)

	//cAgent := new(cag.ConnectionAgent)

	//authService := uaa.NewAuthService(accessService)
	//emailService := et.NewEmailService(requestHandler)

	//policyEnrollment := pc.GetPolicyEnrollmentModel(cAgent)
	//controller := c.GetAuthController()

	// Load balancer check
	mainRouter.Methods(http.MethodGet).Path(pathHeartbeat).Name(routeNameHeartbeat).HandlerFunc(c.SayOK)

	// Route  /ms-login
	mainRouter.Methods(http.MethodGet).Path(pathPrefixAppName).Name(routeNameHeartbeat).HandlerFunc(c.SayOK)

	// Load balancer error check
	mainRouter.Methods(http.MethodGet).Path(pathSayError).Name(routeNameSayError).HandlerFunc(c.SayError)

	return mainRouter

}
