package file_utils

import (
	"bytes"
	"io"
	"reflect"
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
