package main

import (
	"fyne.io/fyne"
	"log"
	"time"
)

type Controller struct {
	app              fyne.App
	model            Model
	preferences      Preferences
	view             View
	ticker           *time.Ticker
	endTime          time.Time
	remainingTime    time.Duration
	roundDone        <-chan bool
	pauseTickingRead <-chan bool
	pauseTickingSend chan<- bool
}

func (ctrl *Controller) startTicking() {
	ctrl.ticker = time.NewTicker(time.Second)
	for {
		select {
		case <-ctrl.pauseTickingRead:
			log.Println("Ticking stopped")
			return
		case <-ctrl.ticker.C:
			remaining := time.Until(ctrl.endTime)
			ctrl.remainingTime = remaining
			log.Println("Tick " + durationToString(remaining))
			if remaining < 0 {
				ctrl.ticker.Stop()
				ctrl.model.runState = PAUSED
				ctrl.goToNextStep()
			} else {
				ctrl.view.setTime(remaining)
			}
		}
	}
}

func (ctrl *Controller) startPauseTimer() {
	if ctrl.model.runState == PAUSED {
		ctrl.model.runState = RUNNING
		if ctrl.remainingTime > 0 {
			// we will continue where we paused
			ctrl.endTime = time.Now().Add(ctrl.remainingTime)
		} else {
			// we finished the step, let's continue with whole step duration
			ctrl.endTime = time.Now().Add(ctrl.model.currentStep.duration)
		}
		ctrl.view.setPause()
		go ctrl.startTicking()
	} else {
		ctrl.model.runState = PAUSED
		ctrl.pauseTickingSend <- true
		ctrl.ticker.Stop()
		ctrl.view.setPlay()
		ctrl.remainingTime = time.Until(ctrl.endTime)
	}
}

func (ctrl *Controller) goToNextStep() {
	ctrl.setNextStep()
	ctrl.view.setPlay()
	ctrl.view.setTime(ctrl.model.currentStep.duration)
	ctrl.view.setRounds(ctrl.model.currentRound, ctrl.preferences.numberOfRounds)
	ctrl.app.SendNotification(
		&fyne.Notification{Title: "Simple Pomodoro", Content: "Step has ended!"},
	)
}

func (ctrl *Controller) setNextStep() {
	nextRound := ctrl.model.findNextRound()
	nextStep := ctrl.model.findNextStep(&ctrl.preferences)
	ctrl.model.currentRound = nextRound
	ctrl.model.currentStep = nextStep
}

func (ctrl *Controller) bindView() {
	ctrl.view.preferences.preferencesChanged = ctrl.savePreferences
	ctrl.view.startPauseTapped = ctrl.startPauseTimer
}

func (ctrl *Controller) savePreferences(rounds int, work int, shortBreak int, longBreak int) {
	pref := ctrl.app.Preferences()
	pref.SetInt("numberOfRounds", rounds)
	pref.SetInt("workDur", work)
	pref.SetInt("breakDur", shortBreak)
	pref.SetInt("longBreakDur", longBreak)
	ctrl.loadPreferences()
	ctrl.view.setRounds(ctrl.model.currentRound, ctrl.preferences.numberOfRounds)
	if ctrl.model.runState == PAUSED {
		ctrl.view.setTime(ctrl.model.currentStep.duration)
	}
}

func (ctrl *Controller) loadPreferences() {
	unit := time.Minute
	//unit := time.Second
	pref := ctrl.app.Preferences()

	workDur := unit * time.Duration(pref.IntWithFallback("workDur", 25))
	breakDur := unit * time.Duration(pref.IntWithFallback("breakDur", 5))
	longBreakDur := unit * time.Duration(pref.IntWithFallback("longBreakDur", 15))

	ctrl.preferences.workStep = Step{kind: WORK, duration: workDur}
	ctrl.preferences.breakStep = Step{kind: BREAK, duration: breakDur}
	ctrl.preferences.longBreakStep = Step{kind: LONG_BREAK, duration: longBreakDur}
	ctrl.preferences.numberOfRounds = pref.IntWithFallback("numberOfRounds", 4)

	if ctrl.model.currentStep.kind == WORK {
		ctrl.model.currentStep = ctrl.preferences.workStep
	} else if ctrl.model.currentStep.kind == BREAK {
		ctrl.model.currentStep = ctrl.preferences.breakStep
	} else {
		ctrl.model.currentStep = ctrl.preferences.longBreakStep
	}
}

func (ctrl *Controller) loadModel() {
	ctrl.model = Model{
		runState:     PAUSED,
		currentRound: 1,
		currentStep:  ctrl.preferences.workStep,
	}
}

func (ctrl *Controller) showApp() {
	ctrl.view.create(ctrl.app, ctrl.model)
	ctrl.view.setRounds(ctrl.model.currentRound, ctrl.preferences.numberOfRounds)
}
