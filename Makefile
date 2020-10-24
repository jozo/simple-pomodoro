pack:
	make clean
	make build


build:
	go build -ldflags "-s -w"
	fyne package -icon icons/app-icon.png -name "Simple Pomodoro" -release -appID "io.jozo.simple-pomodoro"


clean:
	rm -r "Simple Pomodoro.app/"
	rm -f "simple-pomodoro"
