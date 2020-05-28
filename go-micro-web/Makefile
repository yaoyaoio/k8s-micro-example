# author:yao~nfs
GOCMD=go
GOBUILD= $(GOCMD) build

all: build

build: clear build-apline docker-build docker-push

clean: clear

# linux cmd
clear:
	rm -rf web main

build-apline:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD)  -o web main.go

docker-build:
	docker build -t liuyao/go-micro-web:kubernetes .

docker-push:
	docker push liuyao/go-micro-web:kubernetes
