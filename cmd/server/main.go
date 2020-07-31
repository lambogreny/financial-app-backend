package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/namsral/flag"
	"github.com/sirupsen/logrus"

	"9phum.com/financial-app-backend/internal/api"
	"9phum.com/financial-app-backend/internal/config"
	"9phum.com/financial-app-backend/internal/database"
)

//Create Server object and start listerner
func main() {
	flag.Parse()

	logrus.SetLevel(logrus.DebugLevel)

	logrus.WithField("version", config.Version).Debug("starting server.")

	//Creating new database
	db, err := database.New()
	if err != nil {
		logrus.WithError(err).Fatal("Eror verifying database.")
	}

	//Creating new router
	router, err := api.NewRouter(db)
	if err != nil {
		logrus.WithError(err).Fatal("Error building router")
	}

	const addr = "0.0.0.0:8088"
	server := http.Server{
		Handler: router,
		Addr:    addr,
	}

	//Starting server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Error("server failed.")
	}
}
