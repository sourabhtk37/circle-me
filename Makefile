all: build


build:
	go build github.com/turtledev/circle-me

test:
	go test github.com/turtledev/circle-me/...
