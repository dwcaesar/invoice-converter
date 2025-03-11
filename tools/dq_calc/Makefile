.DEFAULT_GOAL := build

clean:
	go clean
	rm builds/* 
build:
	go clean
	env GOOS=linux GOARCH=amd64 go build -o builds/dqLinux main.go
	env GOOS=windows GOARCH=amd64 go build -o builds/dqWin main.go
	env GOOS=linux GOARCH=arm64 go build  -o builds/dqMac main.go
test:
	go test -v