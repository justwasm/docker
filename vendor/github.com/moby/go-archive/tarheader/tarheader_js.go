//go:build js && wasm
// +build js,wasm

package tarheader

import (
	"archive/tar"
	"os"
)

func sysStat(fi os.FileInfo, hdr *tar.Header) error {
	return nil
}
