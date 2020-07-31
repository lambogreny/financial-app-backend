package main

import (
	"9phum.com/financial-app-backend/internal/api"
	"9phum.com/financial-app-backend/internal/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

//Create Server object and start listerner
func main() {

	logrus.SetLevel(logrus.DebugLevel)

	logrus.WithField("version", config.Version).Debug("starting server.")

	router, err := api.NewRouter()
	if err != nil {
		logrus.WithError(err).Fatal("Error building router")
	}

	const addr = "0.0.0.0:8088"
	server := http.Server{
		Handler: router,
		Addr:    addr,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Error("server failed.")
	}
}
