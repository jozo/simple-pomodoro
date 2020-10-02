package main

type RunState int

const (
	Running RunState = iota
	Pause
	Stop
)

type settings struct {
	numberOfRounds int
	interval       int // minutes
	shortBreak     int
	longerBreak    int
}

type model struct {
	runState     RunState
	currentRound int
	settings     settings
}

func loadSettings() settings {
	return settings{
		numberOfRounds: 4,
		interval:       25,
		shortBreak:     5,
		longerBreak:    15,
	}
}

func loadModel() model {
	return model{
		runState:     Stop,
		currentRound: 1,
		settings:     loadSettings(),
	}
}
