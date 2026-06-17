//go:build js && wasm
// +build js,wasm

package user

import (
	"os"
	"path/filepath"
)

func mkdirAs(path string, mode os.FileMode, uid, gid int, mkAll, onlyNew bool) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	if mkAll {
		return os.MkdirAll(path, mode)
	}
	return os.Mkdir(path, mode)
}

func setPermissions(p string, mode os.FileMode, uid, gid int, stat os.FileInfo) error {
	if stat == nil {
		var err error
		stat, err = os.Stat(p)
		if err != nil {
			return err
		}
	}
	if stat.Mode().Perm() != mode.Perm() {
		if err := os.Chmod(p, mode.Perm()); err != nil {
			return err
		}
	}
	return nil
}

func LoadIdentityMapping(name string) (IdentityMapping, error) {
	return IdentityMapping{}, nil
}
