package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"image/color"
	"math"
	"strconv"
	"time"
)

var playIcon, pauseIcon, settingsIcon *theme.ThemedResource

type tappableIcon struct {
	widget.Icon
	OnTapped func()
}

func newTappableIcon(res fyne.Resource, tapped func()) *tappableIcon {
	icon := &tappableIcon{
		OnTapped: tapped,
	}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)

	return icon
}

func (t *tappableIcon) Tapped(_ *fyne.PointEvent) {
	t.OnTapped()
}

func (t *tappableIcon) TappedSecondary(_ *fyne.PointEvent) {
}

func (ctrl *controller) createUI(app fyne.App) {
	w := app.NewWindow("Simple pomodoro")
	playIcon = theme.NewThemedResource(playIconRaw, nil)
	pauseIcon = theme.NewThemedResource(pauseIconRaw, nil)
	settingsIcon = theme.NewThemedResource(settingsIconRaw, nil)
	ctrl.view.timeLabel = canvas.NewText(durationToString(ctrl.model.currentStep.duration), color.White)
	ctrl.view.timeLabel.TextSize = 40
	ctrl.view.roundsLabel = widget.NewLabel(
		fmt.Sprintf("%d/%d",
			ctrl.model.currentRound,
			ctrl.model.settings.numberOfRounds),
	)
	ctrl.view.startPauseButton = widget.NewButtonWithIcon("", playIcon, func() {
		ctrl.startPauseTimer()
	})

	l := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(
			layout.NewCenterLayout(),
			fyne.NewContainerWithLayout(
				layout.NewVBoxLayout(),
				fyne.NewContainerWithLayout(
					layout.NewCenterLayout(),
					ctrl.view.timeLabel,
				),
				fyne.NewContainerWithLayout(
					layout.NewCenterLayout(),
					ctrl.view.startPauseButton,
				),
			),
		),
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			layout.NewSpacer(),
			fyne.NewContainerWithLayout(
				layout.NewCenterLayout(),
				ctrl.view.roundsLabel,
			),
			newTappableIcon(settingsIcon, func() {
				showSettings(app, ctrl.model.settings)
			}),
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

func showSettings(app fyne.App, set Settings) {
	w := app.NewWindow("Preferences")
	f := widget.NewForm(
		widget.NewFormItem("Work (mins)", newEntryWithDuration(set.workStep.duration)),
		widget.NewFormItem("Break (mins)", newEntryWithDuration(set.breakStep.duration)),
		widget.NewFormItem("Long Break (mins)", newEntryWithDuration(set.longBreakStep.duration)),
		widget.NewFormItem("Number of rounds", newEntryWithText(strconv.Itoa(set.numberOfRounds))),
	)
	l := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		fyne.NewContainerWithLayout(
			layout.NewMaxLayout(),
			f,
		),
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(
			layout.NewCenterLayout(),
			widget.NewButton("Save", func() {}),
		),
		layout.NewSpacer(),
	)
	w.Resize(fyne.NewSize(250, 230))
	w.SetContent(l)
	w.Show()
}

func newEntryWithDuration(dur time.Duration) *widget.Entry {
	min := dur / time.Minute
	text := fmt.Sprintf("%d", min)
	return newEntryWithText(text)
}

func newEntryWithText(text string) *widget.Entry {
	e := widget.NewEntry()
	e.SetText(text)
	return e
}

func durationToString(duration time.Duration) string {
	min := duration / time.Minute
	sec := duration - min*time.Minute
	roundedSec := int(math.Round(sec.Seconds()))
	return fmt.Sprintf("%02d:%02d", min, roundedSec)
}
