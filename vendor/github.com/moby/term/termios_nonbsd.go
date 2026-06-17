//go:build !darwin && !freebsd && !netbsd && !openbsd && !windows && !js
// +build !darwin,!freebsd,!netbsd,!openbsd,!windows,!js

package term

import (
	"golang.org/x/sys/unix"
)

const (
	getTermios = unix.TCGETS
	setTermios = unix.TCSETS
)
