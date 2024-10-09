package completion

import (
	"bytes"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

// CobraCompletionCmd represents the completion command
var CobraCompletionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

	source <({{.Use}} completion bash)
	
	# To load completions for each session, execute once:
	# Linux:
	{{.Use}} completion bash > /etc/bash_completion.d/{{.Use}}
	# macOS:
	{{.Use}} completion bash > /usr/local/etc/bash_completion.d/{{.Use}}

Zsh:

	# If shell completion is not already enabled in your environment,
	# you will need to enable it.  You can execute the following once:
	echo "autoload -U compinit; compinit" >> ~/.zshrc

	# To load completions for each session, execute once:
	{{.Use}} completion zsh > "${fpath[1]}/{{.Use}}"

	# You will need to start a new shell for this setup to take effect.

fish:

	{{.Use}} completion fish | source
	
	# To load completions for each session, execute once:
	{{.Use}} completion fish > ~/.config/fish/completions/{{.Use}}.fish

PowerShell:

	{{.Use}} completion powershell | Out-String | Invoke-Expression
	
	# To load completions for every new session, run:
	{{.Use}} completion powershell > {{.Use}}.ps1
	# and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func AddToRootCmd(rootCmd *cobra.Command) {

	tmpl, err := template.New("completion").Parse(CobraCompletionCmd.Long)
	if err != nil {
		panic(err)
	}

	var templateResults bytes.Buffer

	err = tmpl.Execute(&templateResults, rootCmd)
	if err != nil {
		panic(err)
	}

	CobraCompletionCmd.Long = templateResults.String()

	rootCmd.AddCommand(CobraCompletionCmd)
}
