package cobra_updater

import (
	"bytes"
	"github.com/analog-substance/util/cli/updater"
	"github.com/spf13/cobra"
	"text/template"
)

// CobraUpdateCmd represents the image command
var CobraUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update {{.Use}} to latest version",
	Long: `Update or check for updates.
The default update method is to download the latest release from GitHub.

Example 1: Update to latest GitHub

	{{.Use}} update

Example 2: Use go install to update

	{{.Use}} update -g


Example 3: Download from a specific URL

Not sure why anyone else would need this. I use it for quickly testing builds
on different machines.

	{{.Use}} update -u http://10.0.0.2:8000/dist/carbon_darwin_arm64/carbon

This is typically used after I run the following:

	goreleaser release --clean --snapshot
	python -m http.server

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

		updater.SelfUpdate(flag, url)
	},
}

func init() {
	CobraUpdateCmd.Flags().BoolP("check", "C", false, "Check for update")
	CobraUpdateCmd.Flags().BoolP("go-install", "g", false, "Use go install instead of downloading release from GitHub")
	CobraUpdateCmd.Flags().StringP("url", "u", "", "URL to download from (force implies)")
	CobraUpdateCmd.Flags().BoolP("force", "f", false, "Force update, even if release is not newer")
}

func AddToRootCmd(rootCmd *cobra.Command) {

	tmpl, err := template.New("completion").Parse(CobraUpdateCmd.Long)
	if err != nil {
		panic(err)
	}

	var templateResults bytes.Buffer

	err = tmpl.Execute(&templateResults, rootCmd)
	if err != nil {
		panic(err)
	}

	CobraUpdateCmd.Long = templateResults.String()

	rootCmd.AddCommand(CobraUpdateCmd)
}
