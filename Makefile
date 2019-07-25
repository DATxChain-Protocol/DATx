# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gdatx android ios gdatx-cross swarm evm all test clean
.PHONY: gdatx-linux gdatx-linux-386 gdatx-linux-amd64 gdatx-linux-mips64 gdatx-linux-mips64le
.PHONY: gdatx-linux-arm gdatx-linux-arm-5 gdatx-linux-arm-6 gdatx-linux-arm-7 gdatx-linux-arm64
.PHONY: gdatx-darwin gdatx-darwin-386 gdatx-darwin-amd64
.PHONY: gdatx-windows gdatx-windows-386 gdatx-windows-amd64

GOBIN = $(shell pwd)/build/bin
GO ?= latest

gdatx:
	build/env.sh go run build/ci.go install ./cmd/gdatx
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gdatx\" to launch gdatx."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gdatx.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Gdatx.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/jteeuwen/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go install ./cmd/abigen

# Cross Compilation Targets (xgo)

gdatx-cross: gdatx-linux gdatx-darwin gdatx-windows gdatx-android gdatx-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-*

gdatx-linux: gdatx-linux-386 gdatx-linux-amd64 gdatx-linux-arm gdatx-linux-mips64 gdatx-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-linux-*

gdatx-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gdatx
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-linux-* | grep 386

gdatx-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gdatx
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-linux-* | grep amd64

gdatx-linux-arm: gdatx-linux-arm-5 gdatx-linux-arm-6 gdatx-linux-arm-7 gdatx-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-linux-* | grep arm

gdatx-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/gdatx
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-linux-* | grep arm-5

gdatx-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/gdatx
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-linux-* | grep arm-6

gdatx-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/gdatx
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-linux-* | grep arm-7

gdatx-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/gdatx
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-linux-* | grep arm64

gdatx-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gdatx
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-linux-* | grep mips

gdatx-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gdatx
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-linux-* | grep mipsle

gdatx-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gdatx
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-linux-* | grep mips64

gdatx-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gdatx
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-linux-* | grep mips64le

gdatx-darwin: gdatx-darwin-386 gdatx-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-darwin-*

gdatx-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gdatx
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-darwin-* | grep 386

gdatx-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gdatx
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-darwin-* | grep amd64

gdatx-windows: gdatx-windows-386 gdatx-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-windows-*

gdatx-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/gdatx
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-windows-* | grep 386

gdatx-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gdatx
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gdatx-windows-* | grep amd64
