package app

import (
	"fmt"
	"log"
	"net/http"
	"ocp-client/util/config"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type App struct {
	handler http.Handler
	Logger  *logrus.Logger
	*mux.Router
}

func NewApp() *App {
	app := new(App)
	app.initConfig()
	return app
}

func (app *App) Start() {
	fmt.Println("Start")
	addr := app.ServerURL()
	if addr == "" {
		log.Fatal("app.server config missed")
	}
	fmt.Println("Server started at %s.", addr)
	log.Fatal(http.ListenAndServe(addr, app.handler))
}

func (app *App) Stop() {
	fmt.Println("Stop")
}

func (app *App) ServerURL() string {
	return config.GetString("app.server")
}
