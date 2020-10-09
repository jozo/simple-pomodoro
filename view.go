package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"image/color"
	"math"
	"time"
)

func (ctrl *controller) createUI(app fyne.App) {
	w := app.NewWindow("Simple pomodoro")
	ctrl.view.timeLabel = canvas.NewText(durationToString(ctrl.model.currentStep.duration), color.White)
	ctrl.view.timeLabel.TextSize = 40
	ctrl.view.roundsLabel = widget.NewLabel(
		fmt.Sprintf("%d/%d",
			ctrl.model.currentRound,
			ctrl.model.settings.numberOfRounds),
	)
	ctrl.view.startPauseButton = widget.NewButton("Start", func() {
		ctrl.startPauseTimer()
	})
	l := fyne.NewContainerWithLayout(
		layout.NewCenterLayout(),
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewCenterLayout(),
				ctrl.view.timeLabel,
			),
			fyne.NewContainerWithLayout(
				layout.NewCenterLayout(),
				ctrl.view.roundsLabel,
			),
			fyne.NewContainerWithLayout(
				layout.NewCenterLayout(),
				ctrl.view.startPauseButton,
			),
		),
	)
	w.Resize(fyne.NewSize(250, 250))
	w.SetContent(l)
	w.Show()
}

func (ctrl *controller) refreshUIAfterStep() {
	ctrl.view.roundsLabel.Text = fmt.Sprintf("%d/%d", ctrl.model.currentRound, ctrl.model.settings.numberOfRounds)
	ctrl.view.roundsLabel.Refresh()
	ctrl.view.timeLabel.Refresh()
}

func durationToString(duration time.Duration) string {
	min := duration / time.Minute
	sec := duration - min*time.Minute
	roundedSec := int(math.Round(sec.Seconds()))
	return fmt.Sprintf("%02d:%02d", min, roundedSec)
}
