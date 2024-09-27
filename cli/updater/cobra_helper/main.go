package cobra_helper

import (
	"github.com/analog-substance/util/cli/updater"
	"github.com/spf13/cobra"
)

// CobraUpdateCmd represents the image command
var CobraUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update",
	Long:  `Update`,
	Run: func(cmd *cobra.Command, args []string) {
		check, _ := cmd.Flags().GetBool("check")
		downloadRelease, _ := cmd.Flags().GetBool("download-release")
		force, _ := cmd.Flags().GetBool("force")
		url, _ := cmd.Flags().GetString("url")

		var flag updater.OptionsFlag
		if downloadRelease {
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
	CobraUpdateCmd.Flags().BoolP("download-release", "r", false, "Download release instead of go build")
	CobraUpdateCmd.Flags().StringP("url", "u", "", "URL to download from (force implies)")
	CobraUpdateCmd.Flags().BoolP("force", "f", false, "Force update, even if release is not newer")
}
