package main

import (
	"fyne.io/fyne/app"
)

func main() {
	app_ := app.NewWithID("io.jozo.simple-pomodoro")
	model_ := loadModel()

	ticking := make(chan bool, 1)
	pomodoro := &controller{
		roundDone:       make(chan bool, 1),
		stopTickingRead: ticking,
		stopTickingSend: ticking,
	}
	pomodoro.createUI(app_, model_)

	app_.Run()
}
