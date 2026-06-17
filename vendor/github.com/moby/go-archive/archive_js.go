//go:build js && wasm
// +build js,wasm

package archive

import (
	"archive/tar"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func addLongPathPrefix(srcPath string) string {
	return srcPath
}

func getWalkRoot(srcPath string, include string) string {
	return strings.TrimSuffix(srcPath, string(filepath.Separator)) + string(filepath.Separator) + include
}

func chmodTarEntry(perm os.FileMode) os.FileMode {
	return perm
}

func getInodeFromStat(stat interface{}) (uint64, error) {
	return 0, nil
}

func getFileUIDGID(stat interface{}) (int, int, error) {
	return 0, 0, nil
}

func handleTarTypeBlockCharFifo(hdr *tar.Header, path string) error {
	return nil
}

func handleLChmod(hdr *tar.Header, path string, hdrInfo os.FileInfo) error {
	if hdr.Typeflag == tar.TypeSymlink {
		return nil
	}
	return os.Chmod(path, hdrInfo.Mode())
}

func chtimes(name string, atime time.Time, mtime time.Time) error {
	return os.Chtimes(name, atime, mtime)
}

func lchtimes(name string, atime time.Time, mtime time.Time) error {
	return os.Chtimes(name, atime, mtime)
}

func normalizePath(path string) string {
	return path
}

func checkSystemDriveAndRemoveDriveLetter(path string) (string, error) {
	return path, nil
}

func overrideUmask(newMask int) func() {
	return func() {}
}
