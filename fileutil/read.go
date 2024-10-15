package fileutil

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadLowerLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadLowerLinesReader(file)
}

func ReadLowerLinesReader(file io.Reader) ([]string, error) {
	var lines []string
	err := ReadLowerLineByLineReader(file, func(line string) {
		lines = append(lines, line)
	})
	return lines, err
}

func ReadLowerLineMap(path string) (map[string]bool, error) {
	if !FileExists(path) {
		return map[string]bool{}, nil
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ReadLowerLineMapReader(file)
}

func ReadLowerLineMapReader(file io.Reader) (map[string]bool, error) {
	lines := map[string]bool{}
	err := ReadLowerLineByLineReader(file, func(line string) {
		lines[line] = true
	})

	return lines, err
}

func ReadLowerLineByLine(path string, action func(line string)) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return ReadLowerLineByLineReader(file, action)
}

func ReadLowerLineByLineReader(file io.Reader, action func(line string)) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if line != "" {
			action(line)
		}
	}
	return scanner.Err()
}

func ReadLinesMap(path string) (map[string]bool, error) {
	if !FileExists(path) {
		return map[string]bool{}, nil
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ReadLinesMapReader(file)
}

func ReadLinesMapReader(file io.Reader) (map[string]bool, error) {
	lines := map[string]bool{}

	err := ReadLineByLineReader(file, func(line string) {
		lines[line] = true
	})

	return lines, err
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadLinesReader(file)
}

func ReadLinesReader(file io.Reader) ([]string, error) {
	var lines []string
	err := ReadLineByLineReader(file, func(line string) {
		lines = append(lines, line)
	})
	return lines, err
}

func ReadLineByLine(path string, action func(line string)) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return ReadLineByLineReader(file, action)
}

func ReadLineByLineReader(file io.Reader, action func(line string)) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			action(line)
		}
	}
	return scanner.Err()
}

func ReadLineByLineChan(path string) (chan string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// Not sure if this is the best way to re-use ReadLineByLine
	c := make(chan string)
	go func() {
		defer file.Close()

		for s := range ReadLineByLineChanReader(file) {
			c <- s
		}
		close(c)
	}()

	return c, nil
}

func ReadLineByLineChanReader(r io.Reader) chan string {
	c := make(chan string)
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			c <- scanner.Text()
		}
		close(c)
	}()

	return c
}
