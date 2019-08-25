.PHONY: build clean

export GO111MODULE=on
export GOPROXY=https://goproxy.io

all: build

build: tidy
	go build -o tg-keyword-reply-bot.bin .

tidy:
	go mod tidy

clean:
	rm tg-keyword-reply-bot.bin -Rf