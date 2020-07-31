package api

import (
	"net/http"

	"github.com/gorilla/mux"

	v1 "9phum.com/financial-app-backend/internal/api/v1"
	"9phum.com/financial-app-backend/internal/database"
)

//NewRouter provide a handler Api service.
func NewRouter(db database.Database) (http.Handler, error) {
	router := mux.NewRouter()
	router.HandleFunc("/version", v1.VersionHandler)

	return router, nil

}
