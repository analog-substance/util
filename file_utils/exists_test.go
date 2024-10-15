package file_utils

import "testing"

func TestDirExists(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "directory exists",
			args: args{dir: "."},
			want: true,
		},
		{
			name: "directory does not exists",
			args: args{dir: "this-doesnt-exist"},
			want: false,
		},
		{
			name: "directory is a file",
			args: args{dir: "../README.md"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirExists(tt.args.dir); got != tt.want {
				t.Errorf("DirExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileExists(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "file exists",
			args: args{filename: "../.goreleaser.yaml"},
			want: true,
		},
		{
			name: "file doesn't exists",
			args: args{filename: "pizza.party"},
			want: false,
		},
		{
			name: "file is a dir",
			args: args{filename: "."},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExists(tt.args.filename); got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
