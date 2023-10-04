package main

import "ocp-client/cmd/app"

func main() {
	application := app.NewApp()
	defer application.Stop()
	application.Start()
}
