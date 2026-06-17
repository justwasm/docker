//go:build js && wasm
// +build js,wasm

package term

import (
	"io"
	"os"
)

// State holds the platform-specific state / console mode for the terminal.
type State struct{}

// Winsize represents the size of the terminal window.
type Winsize struct {
	Height uint16
	Width  uint16
	x      uint16
	y      uint16
}

// StdStreams returns the standard streams (stdin, stdout, stderr).
func StdStreams() (stdIn io.ReadCloser, stdOut, stdErr io.Writer) {
	return stdStreams()
}

// GetFdInfo returns the file descriptor for an os.File and indicates whether the file represents a terminal.
func GetFdInfo(in interface{}) (fd uintptr, isTerminal bool) {
	return getFdInfo(in)
}

// GetWinsize returns the window size based on the specified file descriptor.
func GetWinsize(fd uintptr) (*Winsize, error) {
	return getWinsize(fd)
}

// SetWinsize tries to set the specified window size for the specified file descriptor.
func SetWinsize(fd uintptr, ws *Winsize) error {
	return setWinsize(fd, ws)
}

// IsTerminal returns true if the given file descriptor is a terminal.
func IsTerminal(fd uintptr) bool {
	return isTerminal(fd)
}

// RestoreTerminal restores the terminal connected to the given file descriptor to a previous state.
func RestoreTerminal(fd uintptr, state *State) error {
	return restoreTerminal(fd, state)
}

// SaveState saves the state of the terminal connected to the given file descriptor.
func SaveState(fd uintptr) (*State, error) {
	return saveState(fd)
}

// DisableEcho applies the specified state to the terminal connected to the file descriptor, with echo disabled.
func DisableEcho(fd uintptr, state *State) error {
	return disableEcho(fd, state)
}

// SetRawTerminal puts the terminal connected to the given file descriptor into raw mode.
func SetRawTerminal(fd uintptr) (previousState *State, err error) {
	return setRawTerminal(fd)
}

// SetRawTerminalOutput puts the output of terminal connected to the given file descriptor into raw mode.
func SetRawTerminalOutput(fd uintptr) (previousState *State, err error) {
	return setRawTerminalOutput(fd)
}

// MakeRaw puts the terminal connected to the given file descriptor into raw mode.
func MakeRaw(fd uintptr) (previousState *State, err error) {
	return makeRaw(fd)
}

func stdStreams() (stdIn io.ReadCloser, stdOut, stdErr io.Writer) {
	return os.Stdin, os.Stdout, os.Stderr
}

func getFdInfo(in interface{}) (uintptr, bool) {
	if file, ok := in.(*os.File); ok {
		return file.Fd(), false
	}
	return 0, false
}

func getWinsize(fd uintptr) (*Winsize, error) {
	return &Winsize{}, nil
}

func setWinsize(fd uintptr, ws *Winsize) error {
	return nil
}

func isTerminal(fd uintptr) bool {
	return false
}

func restoreTerminal(fd uintptr, state *State) error {
	return nil
}

func saveState(fd uintptr) (*State, error) {
	return nil, nil
}

func disableEcho(fd uintptr, state *State) error {
	return nil
}

func setRawTerminal(fd uintptr) (*State, error) {
	return nil, nil
}

func setRawTerminalOutput(fd uintptr) (*State, error) {
	return nil, nil
}

func makeRaw(fd uintptr) (*State, error) {
	return nil, nil
}
