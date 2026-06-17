//go:build js && wasm
// +build js,wasm

package signals

import "os"

var TerminationSignals = []os.Signal{}
