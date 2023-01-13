package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rakesh-gupta29/golang-server/config"
	"github.com/rakesh-gupta29/golang-server/router"
)

func MountAndRun() error {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime) // create a logger

	app_config, err := config.LoadConfig() // loading app config
	if err != nil {
		logger.Fatal("Error loading config file")
	}
	server := &http.Server{
		Addr:         app_config.Address,
		Handler:      router.Mount(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting the server on %s", app_config.Address)
	startError := server.ListenAndServe()
	if startError != nil {
		logger.Fatal(err)
	}
	return nil
}
