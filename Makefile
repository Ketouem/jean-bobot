.PHONY: build

BUILD_FOLDER := ./build
BIN_NAME := jean-bobot

install:
	glide install

build:
	go build -o $(BUILD_FOLDER)/$(BIN_NAME) main.go

run:
	go run main.go

clean:
	find . -name build -type d | xargs rm -fr
	find . -name vendor -type d | xargs rm -fr
