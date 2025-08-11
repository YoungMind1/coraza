package io

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// FSReadFile wraps fs.ReadFile supporting embedio on windows
func FSReadFile(fsys fs.FS, name string) ([]byte, error) {
	return fs.ReadFile(fsys, strings.ReplaceAll(name, string(filepath.Separator), "/"))
}
