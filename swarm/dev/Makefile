.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/DATx/go-DATx/cmd/swarm
	go build -o bin/gdatx     github.com/DATx/go-DATx/cmd/gdatx
	go build -o bin/bootnode github.com/DATx/go-DATx/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/DATx/go-DATx/swarm/...
