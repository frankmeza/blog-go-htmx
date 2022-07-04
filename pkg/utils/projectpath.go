package utils

import (
	"path/filepath"
	"runtime"
)

var (
	_, rootDir, _, _ = runtime.Caller(0)

	// Root folder of this project
	RootDir = filepath.Join(filepath.Dir(rootDir), "../..")
)

func UseRelativePath(path string) string {
	return RootDir + "/" + path
}
