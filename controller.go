package main

import (
	"fmt"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"time"
)

type controller struct {
	timeLabel       *canvas.Text
	ticker          *time.Ticker
	startStopButton *widget.Button
	roundDone       <-chan bool
	stopTickingRead <-chan bool
	stopTickingSend chan<- bool
}

func (ctrl *controller) startTicking() {
	end := time.Now().Add(25 * time.Minute)
	for {
		select {
		case <-ctrl.stopTickingRead:
			fmt.Println("stop ticking")
			return
		case <-ctrl.ticker.C:
			remaining := time.Until(end)
			m := remaining / time.Minute
			s := (remaining - m*time.Minute) / time.Second
			fmt.Println("tick", m, s)
			ctrl.timeLabel.Text = fmt.Sprintf("%02d:%02d", m, s)
			ctrl.timeLabel.Refresh()
		}
	}
}

func (ctrl *controller) startStopTimer() {
	if ctrl.ticker == nil {
		ctrl.ticker = time.NewTicker(time.Second)
		go ctrl.startTicking()
		ctrl.startStopButton.SetText("Stop")
	} else {
		ctrl.stopTickingSend <- true
		ctrl.ticker.Stop()
		ctrl.ticker = nil
		ctrl.timeLabel.Text = "25:00"
		ctrl.timeLabel.Refresh()
		ctrl.startStopButton.SetText("Start")
	}
}
