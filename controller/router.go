package controller

import (
	"fmt"
	"net/http"
	"ocp-client/util/config"

	"github.com/gorilla/mux"
)

func RegisterRouter(router *mux.Router) error {
	installerRouter := router.PathPrefix(config.GetString("router.v1_prefix")).Subrouter()
	installerRouter.HandleFunc(config.GetString("router.provsioning"), Exec).Methods(http.MethodPost)
	fmt.Printf("router: %s%s\n", config.GetString("router.v1_prefix"), config.GetString("router.provsioning"))
	return nil
}
