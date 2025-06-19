// Package main is the entrypoint for the ITSM-ITOM AI Assistant application.
package main

import (
	"fmt"
	"log"
	"net/http"

	"itsm-itom-ai-assistant/config"
	"itsm-itom-ai-assistant/resource"
	"itsm-itom-ai-assistant/restapiv1/routes"
)

// main is the application entrypoint.
func main() {
	// Get the value of the environment variable
	var err error
	config.Conf, err = config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := resource.Connect()
	defer db.Close()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Conf.Port),
		Handler: routes.InitializeRoutes(db),
	}

	fmt.Printf(`server is running on port : %s`, config.Conf.Port)
	log.Fatal(srv.ListenAndServe())
}
