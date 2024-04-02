install:
	go install

build:
	go build -v -o api

start: install build
