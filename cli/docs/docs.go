package docs

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"path"
	"strings"
)

// CobraDocsCmd represents the docs command
var CobraDocsCmd = &cobra.Command{
	Use:    "docs",
	Short:  "Generate docs",
	Long:   `Generate documentation markdown files from the source code.`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {

		outputDir, _ := cmd.Flags().GetString("output-dir")
		identity := func(s string) string { return s }
		emptyStr := func(s string) string {

			caser := cases.Title(language.English)
			title := strings.TrimSuffix(strings.ReplaceAll(path.Base(s), "_", " "), ".md")
			var currentCmd *cobra.Command
			var err error
			if title != cmd.Root().Name() {
				title = strings.TrimSpace(strings.TrimPrefix(title, cmd.Root().Name()))
				args := strings.Split(title, " ")
				currentCmd, _, err = cmd.Root().Find(args)
				if err != nil {
					fmt.Println("error getting command", err)
				}
			} else {
				currentCmd = cmd.Root()
			}

			title = caser.String(title)

			return fmt.Sprintf("---\ntitle: %s\ndescription: %s\n---\n\n", title, currentCmd.Short)
		}

		err := doc.GenMarkdownTreeCustom(cmd.Root(), outputDir, emptyStr, identity)
		//err := doc.GenMarkdownTree(RootCmd, "./docs/cli/")
		if err != nil {
			fmt.Printf("error generating docs err=%s", err)
		}
	},
}

func init() {
	CobraDocsCmd.Flags().StringP("output-dir", "o", "./docs/docs/cli", "Output location")
}
