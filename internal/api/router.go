package api

import (
	"net/http"

	v1 "9phum.com/financial-app-backend/internal/api/v1"
	"github.com/gorilla/mux"
)

//NewRouter provide a handler Api service.
func NewRouter() (http.Handler, error) {
	router := mux.NewRouter()
	router.HandleFunc("/version", v1.VersionHandler)

	return router, nil

}
