#Reduces size by 30%
FLAGS=-ldflags "-w -s"

GOOS=linux
GOARCH=amd64
CC=go build
TARGET=main.go

all: first

first:TARGET=01.1main.go
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(CC) $(FLAGS) $(TARGET) -o ../bin/01.1main



