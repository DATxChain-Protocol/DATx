.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/DATxChain-Protocol/DATx/cmd/swarm
	go build -o bin/gdatx     github.com/DATxChain-Protocol/DATx/cmd/gdatx
	go build -o bin/bootnode github.com/DATxChain-Protocol/DATx/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/DATxChain-Protocol/DATx/swarm/...
