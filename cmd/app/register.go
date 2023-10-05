package app

import "ocp-client/controller"

func (app *App) initRegisters() []Register {
	var registers = []Register{
		controller.RegisterRouter,
	}
	return registers
}
