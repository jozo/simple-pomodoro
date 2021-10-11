package main

import "time"

// Application state - running or paused
type RunState int

const (
	RUNNING RunState = iota
	PAUSED
)

// Steps in Pomodoro
type StepKind int

const (
	WORK StepKind = iota
	BREAK
	LONG_BREAK
)

// A pomodoro step
type Step struct {
	kind     StepKind
	duration time.Duration
}

// App preferences
type Preferences struct {
	numberOfRounds int
	workStep       Step
	breakStep      Step
	longBreakStep  Step
}

// Main model structure
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
