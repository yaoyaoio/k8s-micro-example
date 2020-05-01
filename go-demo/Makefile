# author: yaoyao
GOCMD=go
GOBUILD= $(GOCMD) build

all: build

build: clear build-apline docker-build docker-push

clean: clear

# linux cmd
clear:
	rm -rf main

build-apline:
	CGO_ENABLED=0 $(GOBUILD) main.go

docker-build:
	docker build -t liuyao/go-micro-demo .

docker-push:
	docker push liuyao/go-micro-demo
