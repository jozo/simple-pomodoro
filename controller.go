package main

import (
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"log"
	"time"
)

type viewItems struct {
	timeLabel        *canvas.Text
	roundsLabel      *widget.Label
	startPauseButton *widget.Button
}

type controller struct {
	model            Model
	view             viewItems
	ticker           *time.Ticker
	roundDone        <-chan bool
	pauseTickingRead <-chan bool
	pauseTickingSend chan<- bool
}

func (ctrl *controller) startTicking() {
	ctrl.ticker = time.NewTicker(time.Second)
	end := time.Now().Add(ctrl.model.currentStep.duration)
	for {
		select {
		case <-ctrl.pauseTickingRead:
			log.Println("Ticking stopped")
			return
		case <-ctrl.ticker.C:
			remaining := time.Until(end)
			log.Println("Tick " + durationToString(remaining))
			if remaining < 0 {
				ctrl.ticker.Stop()
				ctrl.model.runState = PAUSED
				ctrl.goToNextStep()
			} else {
				ctrl.updateTimeLabel(remaining)
			}
		}
	}
}

func (ctrl *controller) startPauseTimer() {
	if ctrl.model.runState == PAUSED {
		ctrl.model.runState = RUNNING
		ctrl.view.startPauseButton.SetText("Pause")
		go ctrl.startTicking()
	} else {
		ctrl.model.runState = PAUSED
		ctrl.pauseTickingSend <- true
		ctrl.ticker.Stop()
		ctrl.view.startPauseButton.SetText("Start")
		ctrl.updateTimeLabel(ctrl.model.currentStep.duration)
	}
}

func (ctrl *controller) updateTimeLabel(remaining time.Duration) {
	ctrl.view.timeLabel.Text = durationToString(remaining)
	ctrl.refreshUIAfterStep()
}

func (ctrl *controller) goToNextStep() {
	ctrl.setNextStep()
	// show next step in UI
	ctrl.view.startPauseButton.SetText("Start")
	ctrl.updateTimeLabel(ctrl.model.currentStep.duration)
}

func (ctrl *controller) setNextStep() {
	if ctrl.model.currentStep.kind == LONG_BREAK {
		ctrl.model.currentRound = 1
	} else if ctrl.model.currentStep.kind == WORK {
		ctrl.model.currentRound++
	}

	if ctrl.model.currentStep.kind == WORK {
		if ctrl.model.currentRound%ctrl.model.settings.numberOfRounds == 0 {
			ctrl.model.currentStep = ctrl.model.settings.longBreakStep
		} else {
			ctrl.model.currentStep = ctrl.model.settings.breakStep
		}
	} else {
		ctrl.model.currentStep = ctrl.model.settings.workStep
	}
}
