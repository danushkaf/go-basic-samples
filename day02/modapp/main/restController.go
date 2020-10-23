package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danushkaf/go-basic-samples/day02/modapp/middleware"
	persistentStore "github.com/danushkaf/go-basic-samples/day02/modapp/persistentstore"
	dummyApplicationStore "github.com/danushkaf/go-basic-samples/day02/modapp/persistentstore/dummy/application"
	application "github.com/danushkaf/go-basic-samples/day02/modapp/services/application"
	applicationIRepo "github.com/danushkaf/go-basic-samples/day02/modapp/services/application/repository"
	"github.com/danushkaf/go-basic-samples/day02/modapp/util"
	"github.com/danushkaf/go-basic-samples/day02/modapp/util/common"
	"github.com/gorilla/mux"
)

// main function to boot up everything
func main() {

	dbType := common.GetProperty("dbType")
	var applicationRepository applicationIRepo.Repository

	switch dbType {
	case persistentStore.Dummy:
		applicationRepository = &dummyApplicationStore.DummyApplicationRepository{}
	}

	applicationService := application.NewService(applicationRepository)
	serverPort := util.GetProperties().MustGetString("server-port")
	middlewareChain := middleware.ChainMiddleware(middleware.WithLogging)
	muxRouter := mux.NewRouter()

	// Application Management API routes.
	muxRouter.HandleFunc("/{partnerid}/applications", middlewareChain(application.CreateApplication(applicationService))).Methods("POST")
	muxRouter.HandleFunc("/{partnerid}/applications", common.SupportOptions).Methods("OPTIONS")
	muxRouter.HandleFunc("/{partnerid}/applications/{id}", common.SupportOptions).Methods("OPTIONS")
	muxRouter.HandleFunc("/{partnerid}/applications/{id}", middlewareChain(application.GetApplication(applicationService))).Methods("GET")
	muxRouter.HandleFunc("/{partnerid}/applications/{id}", middlewareChain(application.DeleteApplication(applicationService))).Methods("DELETE")

	fmt.Println("Started shortcode app on port", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, muxRouter))

}
