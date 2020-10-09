package main

import (
	"fyne.io/fyne/app"
)

func main() {
	app_ := app.NewWithID("io.jozo.simple-pomodoro")

	ticking := make(chan bool, 1)
	pomodoro := &controller{
		model:            loadModel(),
		roundDone:        make(chan bool, 1),
		pauseTickingRead: ticking,
		pauseTickingSend: ticking,
	}
	pomodoro.createUI(app_)

	app_.Run()
}
