package v1

import (
	"encoding/json"
	"net/http"

	"9phum.com/financial-app-backend/internal/config"
	log "github.com/sirupsen/logrus"
)

//API for returning version
//When server start, we set version and than use it if necessary

//Serversion reoresent the server version
type ServerVersion struct {
	Version string `json:"version"`
}

//Marshaled Json
var versionJson []byte

func init() {
	var err error
	versionJson, err = json.Marshal(ServerVersion{
		Version: config.Version,
	})
	if err != nil {
		panic(err)
	}

}

func VersionHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	if _, err := w.Write(versionJson); err != nil {
		log.WithError(err).Debug("Eror writting version.")
	}
}
