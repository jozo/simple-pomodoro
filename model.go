package main

import "time"

// RunState represent application state - running or paused
type RunState int

const (
	RUNNING RunState = iota
	PAUSED
)

// StepKind as in Pomodoro
type StepKind int

const (
	WORK StepKind = iota
	BREAK
	LONG_BREAK
)

// Step is a pomodoro step
type Step struct {
	kind     StepKind
	duration time.Duration
}

// Preferences of the app
type Preferences struct {
	numberOfRounds int
	workStep       Step
	breakStep      Step
	longBreakStep  Step
}

// Model contains state of the app
type Model struct {
	runState     RunState
	currentRound int
	currentStep  Step
}

func (m Model) findNextRound() int {
	if m.currentStep.kind == LONG_BREAK {
		return 1
	} else if m.currentStep.kind == BREAK {
		return m.currentRound + 1
	}
	return m.currentRound
}

func (m Model) findNextStep(pref *Preferences) Step {
	if m.currentStep.kind == WORK {
		if m.currentRound == pref.numberOfRounds {
			return pref.longBreakStep
		} else {
			return pref.breakStep
		}
	} else {
		return pref.workStep
	}
}
