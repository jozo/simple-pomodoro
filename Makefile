pack:
	make clean
	make build
	make fyne-pack

fyne-pack:
	fyne package -icon icons/app-icon.png \
	  -name "Simple Pomodoro" \
	  -appID "io.jozo.simple-pomodoro" \
	  -release 

build:
	go build -ldflags "-s -w"


clean:
	rm -r "Simple Pomodoro.app/"
	rm -f "simple-pomodoro"

cross-windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" go build -ldflags "-s -w"

cross-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC="x86_64-linux-musl-gcc" CXX="x86_64-linux-musl-g++" go build -ldflags "-s -w"
