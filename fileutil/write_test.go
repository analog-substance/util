package fileutil

import (
	"os"
	"testing"
)

func TestWriteLines(t *testing.T) {
	type args struct {
		path  string
		lines []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"writes lines",
			args{
				path:  "/tmp/test",
				lines: []string{"one", "two"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteLines(tt.args.path, tt.args.lines); (err != nil) != tt.wantErr {
				t.Errorf("WriteLines() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		err := os.Remove(tt.args.path)
		if err != nil {
			t.Fatal(err)
		}
	}
}
