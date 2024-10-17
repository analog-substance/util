package build_info

import (
	"fmt"
	"os"
	"runtime/debug"
)

type Version struct {
	Version string
	Type    string
	Extra   string
}

func (v Version) String() string {
	return fmt.Sprintf("%s (%s@%s)", v.Version, v.Type, v.Extra)
}

func GetVersion(fallbackVersion, fallbackExtraDetails string) Version {
	buildInfo, ok := debug.ReadBuildInfo()
	buildType := "unknown"
	if ok {
		if fallbackVersion != "v0.0.0" {
			// goreleaser must have set the version
			// lets add gh to the end so we know this release came from github
			buildType = "release"
		} else {
			// not a goreleaser build. lets grab build info from build settings
			fallbackVersion = buildInfo.Main.Version

			if buildInfo.Main.Version == "(devel)" {
				for _, bv := range buildInfo.Settings {
					if bv.Key == "vcs.revision" {
						fallbackExtraDetails = bv.Value[0:8]
						buildType = "go-local"
						break
					}
				}
			} else {
				buildType = "go-remote"
				fallbackExtraDetails = buildInfo.Main.Version
			}
		}
	}

	if os.Getenv("DEBUG_BUILD_INFO") == "1" {
		fmt.Println(buildInfo)
	}

	return Version{
		Version: fallbackVersion,
		Type:    buildType,
		Extra:   fallbackExtraDetails,
	}
}
