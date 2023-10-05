package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Register func(router *mux.Router) error

func (app *App) InitRouter(registers []Register) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homePage).Methods(http.MethodGet)
	r.HandleFunc("/health-check", healthCheck).Methods(http.MethodGet)
	for _, register := range registers {
		fmt.Printf("register: %v\n", register)
		if err := register(r); err != nil {
			fmt.Printf("Register %v error: %v", register, err)
		}
	}
	return r
}

func homePage(w http.ResponseWriter, r *http.Request) {
	result := struct{}{}
	json.NewEncoder(w).Encode(result)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
