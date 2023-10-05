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
	Logger *logrus.Logger
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
	fmt.Printf("Server started at %s.", addr)
	registers := app.initRegisters()
	router := app.InitRouter(registers)
	log.Fatal(http.ListenAndServe(addr, router))
}

func (app *App) Stop() {
	fmt.Println("Stop")
}

func (app *App) ServerURL() string {
	return config.GetString("app.server")
}
