package main

import (
	"fyne.io/fyne/app"
)

func main() {
	app_ := app.NewWithID("io.jozo.simple-pomodoro")

	ticking := make(chan bool, 1)
	controller := &Controller{
		app:              app_,
		roundDone:        make(chan bool, 1),
		pauseTickingRead: ticking,
		pauseTickingSend: ticking,
	}
	controller.loadPreferences()
	controller.loadModel()
	controller.bindView()
	controller.showApp()

	app_.Run()
}
