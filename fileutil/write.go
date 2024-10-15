package fileutil

import (
	"github.com/analog-substance/util/string_utils"
	"io/fs"
	"os"
	"strings"
)

const (
	DefaultDirPerms  fs.FileMode = 0755
	DefaultFilePerms fs.FileMode = 0644
)

func WriteLines(path string, lines []string) error {
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), DefaultFilePerms)
}

func WriteLowerUniqueLines(path string, lines []string) error {
	sortedUnique := string_utils.SortedLowerUnique(lines)
	return os.WriteFile(path, []byte(strings.Join(sortedUnique, "\n")+"\n"), DefaultFilePerms)
}

func WriteString(path string, content string) error {
	return os.WriteFile(path, []byte(content), DefaultFilePerms)
}
