//go:build !windows

package io

import (
	"io/fs"
)

// FSReadFile wraps fs.ReadFile supporting embedio on windows
var FSReadFile = fs.ReadFile
