//go:build js && wasm
// +build js,wasm

package signal

import "syscall"

const (
	SIGCHLD  = syscall.Signal(0)
	SIGWINCH = syscall.Signal(0)
	SIGPIPE  = syscall.Signal(0)
)
