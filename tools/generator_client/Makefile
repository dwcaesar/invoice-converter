.DEFAULT_GOAL := build

clean:
	go clean
	rm builds/*
build:
	go clean
	env GOOS=linux GOARCH=amd64 go build -o builds/gcLinux main.go
	env GOOS=windows GOARCH=amd64 go build -o builds/gcWin main.go
	env GOOS=linux GOARCH=arm64 go build  -o builds/gcMac main.go
test:
	go test -v