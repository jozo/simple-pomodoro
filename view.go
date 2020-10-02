package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"image/color"
	"strconv"
)

func (ctrl *controller) createUI(app fyne.App, appModel model) {
	w := app.NewWindow("Simple pomodoro")
	timeLbl := canvas.NewText(strconv.Itoa(appModel.settings.interval)+":00", color.White)
	timeLbl.TextSize = 40
	startBtn := widget.NewButton("Start", func() {
		ctrl.startStopTimer()
	})
	ctrl.timeLabel = timeLbl
	ctrl.startStopButton = startBtn
	l := fyne.NewContainerWithLayout(
		layout.NewCenterLayout(),
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			fyne.NewContainerWithLayout(
				layout.NewCenterLayout(),
				timeLbl,
			),
			fyne.NewContainerWithLayout(
				layout.NewCenterLayout(),
				startBtn,
			),
		),
	)
	w.Resize(fyne.NewSize(250, 250))
	w.SetContent(l)
	w.Show()
}
