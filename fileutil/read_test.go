package fileutil

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestReadLinesMapReader(t *testing.T) {
	type args struct {
		file io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]bool
		wantErr bool
	}{
		{
			"read",
			args{file: bytes.NewReader([]byte("hello\nworld"))},
			map[string]bool{"hello": true, "world": true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadLinesMapReader(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLinesMapReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadLinesMapReader() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadLinesMap(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]bool
		wantErr bool
	}{
		{
			"no new line",
			args{path: "test-no-newline.txt"},
			map[string]bool{"One": true, "Two": true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadLinesMap(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLinesMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadLinesMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadLowerLines(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"no new line",
			args{path: "test-no-newline.txt"},
			[]string{"one", "two"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadLowerLines(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLowerLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadLowerLines() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadLowerLineMap(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]bool
		wantErr bool
	}{
		{
			"no new line",
			args{path: "test-no-newline.txt"},
			map[string]bool{"one": true, "two": true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadLowerLineMap(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLowerLineMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadLowerLineMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadLines(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"no new line",
			args{path: "test-no-newline.txt"},
			[]string{"One", "Two"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadLines(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadLines() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadFileLineByLine(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Multi line",
			args: args{
				r: strings.NewReader("line 1\nline 2\nline 3"),
			},
			want: []string{
				"line 1",
				"line 2",
				"line 3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChan := ReadLineByLineChanReader(tt.args.r)

			var got []string
			for g := range gotChan {
				got = append(got, g)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFileLineByLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadLineByLine(t *testing.T) {
	tempDir := t.TempDir()

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		setup   func()
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Multi line file",
			setup: func() {
				os.WriteFile(filepath.Join(tempDir, "test.txt"), []byte("line 1\nline 2\nline 3"), 0644)
			},
			args: args{
				path: filepath.Join(tempDir, "test.txt"),
			},
			want: []string{
				"line 1",
				"line 2",
				"line 3",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			gotChan, err := ReadLineByLineChan(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLineByLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var got []string
			for g := range gotChan {
				got = append(got, g)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadLineByLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
