// Copyright 2016 The go-DATx Authors
// This file is part of go-DATx.
//
// go-DATx is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-DATx is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-DATx. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"crypto/rand"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/DATxChain-Protocol/DATx/params"
)

const (
	ipcAPIs  = "admin:1.0 debug:1.0 dpos:1.0 datx:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 shh:1.0 txpool:1.0 DATxWeb:1.0"
	httpAPIs = "datx:1.0 net:1.0 rpc:1.0 DATxWeb:1.0"
)

// Tests that a node embedded within a console can be started up properly and
// then terminated by closing the input stream.
func TestConsoleWelcome(t *testing.T) {
	validator := "0x34d4a8d9f6b53a8f5e674516cb8ad66c843b2801"
	coinbase := "0x8605cdbbdb6d264aa742e77020dcbc58fcdce182"

	// Start a gdatx console, make sure it's cleaned up and terminate the console
	gdatx := runGdatx(t,
		"--port", "0", "--maxpeers", "0", "--nodiscover", "--nat", "none",
		"--coinbase", coinbase, "--validator", validator, "--shh",
		"console")

	// Gather all the infos the welcome message needs to contain
	gdatx.SetTemplateFunc("goos", func() string { return runtime.GOOS })
	gdatx.SetTemplateFunc("goarch", func() string { return runtime.GOARCH })
	gdatx.SetTemplateFunc("gover", runtime.Version)
	gdatx.SetTemplateFunc("gdatxver", func() string { return params.Version })
	gdatx.SetTemplateFunc("niltime", func() string { return time.Unix(1522052340, 0).Format(time.RFC1123) })
	gdatx.SetTemplateFunc("apis", func() string { return ipcAPIs })

	// Verify the actual welcome message to the required template
	gdatx.Expect(`
Welcome to the Gdatx JavaScript console!

 instance: Gdatx/v{{gdatxver}}/{{goos}}-{{goarch}}/{{gover}}
validator: {{.Validator}}
 coinbase: {{.Coinbase}}
 at block: 0 ({{niltime}})
  datadir: {{.Datadir}}
  modules: {{apis}}

> {{.InputLine "exit"}}
`)
	gdatx.ExpectExit()
}

// Tests that a console can be attached to a running node via various means.
func TestIPCAttachWelcome(t *testing.T) {
	// Configure the instance for IPC attachement
	validator := "0x34d4a8d9f6b53a8f5e674516cb8ad66c843b2801"
	coinbase := "0x8605cdbbdb6d264aa742e77020dcbc58fcdce182"
	var ipc string
	if runtime.GOOS == "windows" {
		ipc = `\\.\pipe\gdatx` + strconv.Itoa(trulyRandInt(100000, 999999))
	} else {
		ws := tmpdir(t)
		defer os.RemoveAll(ws)
		ipc = filepath.Join(ws, "gdatx.ipc")
	}
	// Note: we need --shh because testAttachWelcome checks for default
	// list of ipc modules and shh is included there.
	gdatx := runGdatx(t,
		"--port", "0", "--maxpeers", "0", "--nodiscover", "--nat", "none",
		"--coinbase", coinbase, "--validator", validator, "--shh", "--ipcpath", ipc)

	time.Sleep(2 * time.Second) // Simple way to wait for the RPC endpoint to open
	testAttachWelcome(t, gdatx, "ipc:"+ipc, ipcAPIs)

	gdatx.Interrupt()
	gdatx.ExpectExit()
}

func TestHTTPAttachWelcome(t *testing.T) {
	validator := "0x34d4a8d9f6b53a8f5e674516cb8ad66c843b2801"
	coinbase := "0x8605cdbbdb6d264aa742e77020dcbc58fcdce182"
	port := strconv.Itoa(trulyRandInt(1024, 65536)) // Yeah, sometimes this will fail, sorry :P
	gdatx := runGdatx(t,
		"--port", "0", "--maxpeers", "0", "--nodiscover", "--nat", "none",
		"--coinbase", coinbase, "--validator", validator, "--rpc", "--rpcport", port)

	time.Sleep(2 * time.Second) // Simple way to wait for the RPC endpoint to open
	testAttachWelcome(t, gdatx, "http://localhost:"+port, httpAPIs)

	gdatx.Interrupt()
	gdatx.ExpectExit()
}

func TestWSAttachWelcome(t *testing.T) {
	validator := "0x34d4a8d9f6b53a8f5e674516cb8ad66c843b2801"
	coinbase := "0x8605cdbbdb6d264aa742e77020dcbc58fcdce182"
	port := strconv.Itoa(trulyRandInt(1024, 65536)) // Yeah, sometimes this will fail, sorry :P

	gdatx := runGdatx(t,
		"--port", "0", "--maxpeers", "0", "--nodiscover", "--nat", "none",
		"--coinbase", coinbase, "--validator", validator, "--ws", "--wsport", port)

	time.Sleep(2 * time.Second) // Simple way to wait for the RPC endpoint to open
	testAttachWelcome(t, gdatx, "ws://localhost:"+port, httpAPIs)

	gdatx.Interrupt()
	gdatx.ExpectExit()
}

func testAttachWelcome(t *testing.T, gdatx *testgdatx, endpoint, apis string) {
	// Attach to a running gdatx note and terminate immediately
	attach := runGdatx(t, "attach", endpoint)
	defer attach.ExpectExit()
	attach.CloseStdin()

	// Gather all the infos the welcome message needs to contain
	attach.SetTemplateFunc("goos", func() string { return runtime.GOOS })
	attach.SetTemplateFunc("goarch", func() string { return runtime.GOARCH })
	attach.SetTemplateFunc("gover", runtime.Version)
	attach.SetTemplateFunc("gdatxver", func() string { return params.Version })
	attach.SetTemplateFunc("validator", func() string { return gdatx.Validator })
	attach.SetTemplateFunc("coinbase", func() string { return gdatx.Coinbase })
	attach.SetTemplateFunc("niltime", func() string { return time.Unix(1522052340, 0).Format(time.RFC1123) })
	attach.SetTemplateFunc("ipc", func() bool { return strings.HasPrefix(endpoint, "ipc") })
	attach.SetTemplateFunc("datadir", func() string { return gdatx.Datadir })
	attach.SetTemplateFunc("apis", func() string { return apis })

	// Verify the actual welcome message to the required template
	attach.Expect(`
Welcome to the Gdatx JavaScript console!

 instance: Gdatx/v{{gdatxver}}/{{goos}}-{{goarch}}/{{gover}}
validator: {{validator}}
 coinbase: {{coinbase}}
 at block: 0 ({{niltime}}){{if ipc}}
  datadir: {{datadir}}{{end}}
  modules: {{apis}}

> {{.InputLine "exit" }}
`)
	attach.ExpectExit()
}

// trulyRandInt generates a crypto random integer used by the console tests to
// not clash network ports with other tests running cocurrently.
func trulyRandInt(lo, hi int) int {
	num, _ := rand.Int(rand.Reader, big.NewInt(int64(hi-lo)))
	return int(num.Int64()) + lo
}
