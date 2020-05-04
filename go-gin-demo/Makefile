# author:yao~nfs
GOCMD=go
GOBUILD= $(GOCMD) build

all: build

build: clear build-apline docker-build docker-push

clean: clear

# linux cmd
clear:
	rm -rf main

build-apline:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) main.go

docker-build:
	docker build -t liuyao/go-gin-demo .

docker-push:
	docker push liuyao/go-gin-demo
