# simple-pomodoro

## Icons
https://github.com/tabler/tabler-icons

```shell
fyne bundle -package main -name playIconRaw icons/player-play.svg > icons.go
fyne bundle -package main -name pauseIconRaw -append icons/player-pause.svg >> icons.go
fyne bundle -package main -name settingsIconRaw -append icons/settings.svg >> icons.go
```

## Packaging
```shell
fyne package -icon icons/app-icon.png -name "Simple Pomodoro" -release -appID "io.jozo.simple-pomodoro"
```

## Consider
https://github.com/gen2brain/beeep
