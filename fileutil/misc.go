package fileutil

import (
	"errors"
	"fmt"
	cp "github.com/otiai10/copy"
	"io"
	"os"
	"path/filepath"
)

func IsSameFile(p1 string, p2 string) bool {
	p1Info, err := os.Stat(p1)
	if err != nil {
		return false
	}

	p2Info, err := os.Stat(p2)
	if err != nil {
		return false
	}

	return os.SameFile(p1Info, p2Info)
}

func CopyFile(src string, dest string) error {
	if IsSameFile(src, dest) {
		return errors.New("source and destination are the same file")
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	srcInfo, _ := srcFile.Stat()

	if !srcInfo.Mode().IsRegular() {
		return fmt.Errorf("non-regular source file %s (%q)", srcInfo.Name(), srcInfo.Mode().String())
	}

	var destFile *os.File
	if FileExists(dest) { // dest is path/to/existing/file
		destFile, err = os.OpenFile(dest, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, srcInfo.Mode().Perm())
	} else if DirExists(dest) { // dest is path/to/existing/dir
		destFile, err = os.OpenFile(filepath.Join(dest, srcInfo.Name()), os.O_CREATE|os.O_WRONLY, srcInfo.Mode().Perm())
	} else if DirExists(filepath.Dir(dest)) { // dest is path/to/existing/dir/non_existent_file
		destFile, err = os.OpenFile(filepath.Join(filepath.Dir(dest), filepath.Base(dest)), os.O_CREATE|os.O_WRONLY, srcInfo.Mode().Perm())
	} else { // dest is to a path that doesn't exist
		err = errors.New("destination path doesn't exist")
	}

	if err != nil {
		return err
	}
	defer destFile.Close()

	written, err := io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	if written != srcInfo.Size() {
		return errors.New("error writing data to destination file")
	}

	return nil
}

func CopyDir(src string, dest string) error {
	if !DirExists(src) {
		return errors.New("source path doesn't exist or isn't a directory")
	}

	if IsSameFile(src, dest) {
		return errors.New("source and destination are the same file")
	}

	if DirExists(dest) {
		dest = filepath.Join(dest, filepath.Base(src))
	} else if FileExists(dest) {
		return errors.New("destination is a file")
	}

	return cp.Copy(src, dest, cp.Options{
		PreserveOwner: true,
	})
}

func HasStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	mode := stat.Mode()

	isPipedFromChrDev := (mode & os.ModeCharDevice) == 0
	isPipedFromFIFO := (mode & os.ModeNamedPipe) != 0

	return isPipedFromChrDev || isPipedFromFIFO
}

func MkdirAll(dirs ...string) []error {
	var errors []error
	for _, dir := range dirs {
		err := os.MkdirAll(dir, DefaultDirPerms)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}
