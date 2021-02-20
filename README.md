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
fyne bundle -package main -name alarmSound "sounds/Electronic Beeping Alarm Clock.wav" > sounds.go
```


## Cross-compiling

### To linux
```shell
brew install mesa
brew install --cask xquartz
ln -s /opt/X11/include/X11 /usr/local/include/X11
```
