package cobra_updater

import (
	"bytes"
	"github.com/analog-substance/util/cli/updater"
	"github.com/analog-substance/util/cli/version"
	"github.com/spf13/cobra"
	"text/template"
)

var versionInfo version.Info

// CobraUpdateCmd represents the image command
var CobraUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update {{.Use}} to latest version",
	Long: `Update or check for updates.
The default update method is to download the latest release from GitHub.`,
	Example: `# Update to latest version
{{.Use}} update


# Use go install to update
{{.Use}} update -g


# Download from a specific URL
# Not sure why anyone else would need this. I use it for quickly testing builds on different machines.
{{.Use}} update -u http://10.0.0.2:8000/dist/carbon_darwin_arm64/carbon

# This is typically used after I run the following:
#	goreleaser release --clean --snapshot
#	python -m http.server

`,
	Run: func(cmd *cobra.Command, args []string) {
		check, _ := cmd.Flags().GetBool("check")
		goInstall, _ := cmd.Flags().GetBool("go-install")
		force, _ := cmd.Flags().GetBool("force")
		url, _ := cmd.Flags().GetString("url")

		var flag updater.OptionsFlag
		if !goInstall {
			flag = flag | updater.OptionsRelease
		}
		if force {
			flag = flag | updater.OptionsForce
		}
		if check {
			flag = flag | updater.OptionsCheck
		}

		updater.SelfUpdate(flag, url, versionInfo)
	},
}

func init() {
	CobraUpdateCmd.Flags().BoolP("check", "C", false, "Check for update")
	CobraUpdateCmd.Flags().BoolP("go-install", "g", false, "Use go install instead of downloading release from GitHub")
	CobraUpdateCmd.Flags().StringP("url", "u", "", "URL to download from (force implies)")
	CobraUpdateCmd.Flags().BoolP("force", "f", false, "Force update, even if release is not newer")
}

func AddToRootCmd(rootCmd *cobra.Command, info version.Info) {

	versionInfo = info

	longTmpl, err := template.New("long").Parse(CobraUpdateCmd.Long)
	if err != nil {
		panic(err)
	}
	exampleTmpl, err := template.New("example").Parse(CobraUpdateCmd.Example)
	if err != nil {
		panic(err)
	}
	shortTmpl, err := template.New("short").Parse(CobraUpdateCmd.Short)
	if err != nil {
		panic(err)
	}
	var longTmplResult bytes.Buffer
	var exampleTmplResult bytes.Buffer
	var shortTmplResult bytes.Buffer

	err = longTmpl.Execute(&longTmplResult, rootCmd)
	if err != nil {
		panic(err)
	}

	err = exampleTmpl.Execute(&exampleTmplResult, rootCmd)
	if err != nil {
		panic(err)
	}

	err = shortTmpl.Execute(&shortTmplResult, rootCmd)
	if err != nil {
		panic(err)
	}
	CobraUpdateCmd.Long = longTmplResult.String()
	CobraUpdateCmd.Example = exampleTmplResult.String()
	CobraUpdateCmd.Short = shortTmplResult.String()
	rootCmd.AddCommand(CobraUpdateCmd)

}
