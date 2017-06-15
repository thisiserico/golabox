.PHONY: all build run

all: build run

build:
	go build -o golabox main.go

run:
	./golabox
