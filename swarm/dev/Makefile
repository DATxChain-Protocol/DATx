.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/KunkaYU/go-DATx/cmd/swarm
	go build -o bin/gdatx     github.com/KunkaYU/go-DATx/cmd/gdatx
	go build -o bin/bootnode github.com/KunkaYU/go-DATx/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/KunkaYU/go-DATx/swarm/...
