// +build !windows

package rabbitgod

import (
	"os"
)

func atomicRename(sourceFile, targetFile string) error {
	return os.Rename(sourceFile, targetFile)
}
