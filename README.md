# simple-pomodoro

## Packaging
_Note_: I changed SendNotification in app_darwin.go from fyne to:
```go
template := `display notification "%s" with title "%s" sound name "default"`
```

**TO BUILD**:
```shell
make build
```

# Sounds
https://freesfx.co.uk

# Resources
```shell
fyne bundle -package main -name alarmSound "sounds/Electronic Beeping Alarm Clock.ogg" > sounds.go
```