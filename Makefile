pack:
	make clean
	make build


build:
	go build -ldflags "-s -w"
	fyne package -icon icons/app-icon.png -name "Simple Pomodoro" -release -appID "io.jozo.simple-pomodoro"


clean:
	rm -r "Simple Pomodoro.app/"
	rm -f "simple-pomodoro"

windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" go build -ldflags "-s -w"

linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC="x86_64-linux-musl-gcc" CXX="x86_64-linux-musl-g++" go build -ldflags "-s -w"
