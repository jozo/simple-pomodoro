package main

import "time"

type RunState int

const (
	RUNNING RunState = iota
	PAUSED
)

type StepKind int

const (
	WORK StepKind = iota
	BREAK
	LONG_BREAK
)

type Step struct {
	kind     StepKind
	duration time.Duration
}

type Settings struct {
	numberOfRounds int
	workStep       Step
	breakStep      Step
	longBreakStep  Step
}

type Model struct {
	runState     RunState
	currentRound int
	currentStep  Step
	settings     Settings
}

func loadSettings() Settings {
	return Settings{
		numberOfRounds: 4,
		//workStep:       Step{kind: WORK, duration: 5 * time.Second},
		//breakStep:      Step{kind: BREAK, duration: 2 * time.Second},
		//longBreakStep:  Step{kind: LONG_BREAK, duration: 7 * time.Second},
		workStep:      Step{kind: WORK, duration: 25 * time.Minute},
		breakStep:     Step{kind: BREAK, duration: 5 * time.Minute},
		longBreakStep: Step{kind: LONG_BREAK, duration: 15 * time.Minute},
	}
}

func loadModel() Model {
	settings := loadSettings()
	return Model{
		runState:     PAUSED,
		currentRound: 1,
		currentStep:  settings.workStep,
		settings:     settings,
	}
}
