package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.NewWithID("io.jozo.simple-pomodoro")

	ticking := make(chan bool, 1)
	controller := &Controller{
		app:              app,
		roundDone:        make(chan bool, 1),
		pauseTickingRead: ticking,
		pauseTickingSend: ticking,
	}
	controller.loadPreferences()
	controller.loadModel()
	controller.bindView()
	controller.showApp()

	app.Run()
}
