# author:yao~nfs
GOCMD=go
GOBUILD= $(GOCMD) build

all: build

build: clear build-apline docker-build docker-push

clean: clear

# linux cmd
clear:
	rm -rf server main

build-apline:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD)  -o server main.go

docker-build:
	docker build -t liuyao/go-micro-srv:kubernetes .

docker-push:
	docker push liuyao/go-micro-srv:kubernetes
