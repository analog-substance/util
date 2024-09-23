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
		updater.SelfUpdate()
	},
}
