package updater

import "testing"

func Test_isNewerVersion(t *testing.T) {
	type args struct {
		currentVersion   string
		versionToCompare string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"is newer than v0.0.0",
			args{
				currentVersion:   "v0.0.0",
				versionToCompare: "v0.0.1",
			},
			true,
		},
		{
			"is newer than v0.1.0",
			args{
				currentVersion:   "v0.1.0",
				versionToCompare: "v0.1.1",
			},
			true,
		},
		{
			"is newer than v0.1.9",
			args{
				currentVersion:   "v0.1.9",
				versionToCompare: "v0.1.10",
			},
			true,
		},
		{
			"is newer than v0.1.10",
			args{
				currentVersion:   "v0.1.10",
				versionToCompare: "v0.1.11",
			},
			true,
		},
		{
			"is newer than v0.1.1",
			args{
				currentVersion:   "v0.1.1",
				versionToCompare: "v0.1.11",
			},
			true,
		},
		{
			"is newer than v0.1.45",
			args{
				currentVersion:   "v0.1.45",
				versionToCompare: "v0.2.00",
			},
			true,
		},
		{
			"is not newer than v0.1.40",
			args{
				currentVersion:   "v0.1.40",
				versionToCompare: "v0.1.4",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNewerVersion(tt.args.currentVersion, tt.args.versionToCompare); got != tt.want {
				t.Errorf("isNewerVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
