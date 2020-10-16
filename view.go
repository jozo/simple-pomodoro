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

type View struct {
	preferences      PreferencesView
	roundsLabel      *widget.Label
	timeLabel        *canvas.Text
	startPauseButton *widget.Button

	// callbacks
	startPauseTapped func()
}

func (view *View) create(app fyne.App, model Model) {
	loadIcons()
	w := app.NewWindow("Simple pomodoro")
	view.timeLabel = canvas.NewText(durationToString(model.currentStep.duration), color.White)
	view.timeLabel.TextSize = 40
	view.roundsLabel = widget.NewLabel("")
	view.startPauseButton = widget.NewButtonWithIcon("", playIcon, func() {
		view.startPauseTapped()
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
					view.timeLabel,
				),
				fyne.NewContainerWithLayout(
					layout.NewCenterLayout(),
					view.startPauseButton,
				),
			),
		),
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			layout.NewSpacer(),
			fyne.NewContainerWithLayout(
				layout.NewCenterLayout(),
				view.roundsLabel,
			),
			newTappableIcon(settingsIcon, func() {
				view.preferences.create(app)
			}),
		),
	)
	w.Resize(fyne.NewSize(250, 250))
	w.SetContent(l)
	w.Show()
}

func (view *View) setRounds(round int, rounds int) {
	view.roundsLabel.SetText(fmt.Sprintf("%d/%d", round, rounds))
	view.roundsLabel.Refresh()
}

func (view *View) setPause() {
	view.startPauseButton.SetIcon(pauseIcon)
}

func (view *View) setPlay() {
	view.startPauseButton.SetIcon(playIcon)
}

func (view *View) setTime(remaining time.Duration) {
	view.timeLabel.Text = durationToString(remaining)
	view.timeLabel.Refresh()
}

type PreferencesView struct {
	workEntry      *widget.Entry
	breakEntry     *widget.Entry
	longBreakEntry *widget.Entry
	roundsEntry    *widget.Entry

	// callbacks
	preferencesChanged func(int, int, int, int)
}

func (view PreferencesView) extract() (int, int, int, int) {
	rounds, _ := strconv.Atoi(view.roundsEntry.Text)
	work, _ := strconv.Atoi(view.workEntry.Text)
	break_, _ := strconv.Atoi(view.breakEntry.Text)
	longBreak, _ := strconv.Atoi(view.longBreakEntry.Text)
	return rounds, work, break_, longBreak
}

func (view *PreferencesView) create(app fyne.App) {
	w := app.NewWindow("Preferences")
	pref := app.Preferences()
	view.workEntry = newNumberEntry(pref.IntWithFallback("workDur", 25))
	view.breakEntry = newNumberEntry(pref.IntWithFallback("breakDur", 5))
	view.longBreakEntry = newNumberEntry(pref.IntWithFallback("longBreakDur", 15))
	view.roundsEntry = newNumberEntry(pref.IntWithFallback("numberOfRounds", 4))

	f := widget.NewForm(
		widget.NewFormItem("Work (mins)", view.workEntry),
		widget.NewFormItem("Break (mins)", view.breakEntry),
		widget.NewFormItem("Long Break (mins)", view.longBreakEntry),
		widget.NewFormItem("Number of rounds", view.roundsEntry),
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
			widget.NewButton("Save", func() {
				view.preferencesChanged(view.extract())
				w.Close()
			}),
		),
		layout.NewSpacer(),
	)
	w.Resize(fyne.NewSize(250, 230))
	w.SetContent(l)
	w.Show()
}

func newNumberEntry(number int) *widget.Entry {
	e := widget.NewEntry()
	t := strconv.Itoa(number)
	e.SetText(t)
	e.OnChanged = func(s string) {
		_, err := strconv.Atoi(s)
		if err != nil {
			e.SetText(t)
		}
	}
	return e
}

func durationToString(duration time.Duration) string {
	min := duration / time.Minute
	sec := duration - min*time.Minute
	roundedSec := int(math.Round(sec.Seconds()))
	return fmt.Sprintf("%02d:%02d", min, roundedSec)
}

func loadIcons() {
	playIcon = theme.NewThemedResource(playIconRaw, nil)
	pauseIcon = theme.NewThemedResource(pauseIconRaw, nil)
	settingsIcon = theme.NewThemedResource(settingsIconRaw, nil)
}
