build:
	make clean
	go build -ldflags "-s -w"
	fyne package -icon icons/app-icon.png -name "Simple Pomodoro" -release -appID "io.jozo.simple-pomodoro"


clean:
	rm -rf "Simple Pomodor.app" "simple-pomodoro"
