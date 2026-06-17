//go:build !windows && !js

package archive

// checkSystemDriveAndRemoveDriveLetter is the non-Windows implementation
// of CheckSystemDriveAndRemoveDriveLetter
func checkSystemDriveAndRemoveDriveLetter(path string) (string, error) {
	return path, nil
}
