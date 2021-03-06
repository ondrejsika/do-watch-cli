build:
	go build

install:
	go install

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o dist/do-watch-cli-darwin-amd64

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o dist/do-watch-cli-linux-amd64
