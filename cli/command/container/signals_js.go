//go:build js && wasm
// +build js,wasm

package container

import "os"

func isRuntimeSig(s os.Signal) bool {
	return false
}
