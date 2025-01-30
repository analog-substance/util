package docs

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
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

			isRootCommand := strings.HasSuffix(s, fmt.Sprintf("%c_index.md", os.PathSeparator))

			titleCaser := cases.Title(language.English)
			title := strings.TrimSuffix(strings.ReplaceAll(filepath.Base(s), "_", " "), markdownExtension)

			var prepend string
			if isRootCommand {
				commandNameTitle := titleCaser.String(cmd.Root().Name())
				prepend = fmt.Sprintf("---\ntitle: %s CLI\ndescription: %s CLI Reference\nno_list: true\nweight: 1\n---\n", commandNameTitle, commandNameTitle)

			} else {
				title = strings.TrimSpace(strings.TrimPrefix(title, cmd.Root().Name()))
				args := strings.Split(title, " ")
				currentCmd, _, err := cmd.Root().Find(args)
				if err != nil {
					fmt.Println("error getting command", err)
				}
				title = titleCaser.String(title)
				prepend = fmt.Sprintf("---\ntitle: %s\ndescription: %s\n---\n\n", title, currentCmd.Short)
			}

			return prepend
		}

		err := GenMarkdownTreeCustom(cmd.Root(), outputDir, emptyStr, identity)
		if err != nil {
			fmt.Printf("error generating docs err=%s", err)
		}
	},
}

func init() {
	outputPath := filepath.Join(".", "docs", "docs", "cli")
	CobraDocsCmd.Flags().StringP("output-dir", "o", outputPath, "Output location")
}

const markdownExtension = ".md"

// GenMarkdownTreeCustom is the same as GenMarkdownTree, but
// with custom filePrepender and linkHandler.
func GenMarkdownTreeCustom(cmd *cobra.Command, dir string, filePrepender, linkHandler func(string) string) error {
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		if err := GenMarkdownTreeCustom(c, dir, filePrepender, linkHandler); err != nil {
			return err
		}
	}

	basename := strings.ReplaceAll(cmd.CommandPath(), " ", "_") + markdownExtension

	if cmd.Parent() == nil {
		// we are the root command
		basename = "_index" + markdownExtension
	}
	filename := filepath.Join(dir, basename)

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.WriteString(f, filePrepender(filename)); err != nil {
		return err
	}
	if err := GenMarkdownCustom(cmd, f, linkHandler); err != nil {
		return err
	}
	return nil
}

// GenMarkdownCustom creates custom markdown output.
func GenMarkdownCustom(cmd *cobra.Command, w io.Writer, linkHandler func(string) string) error {
	cmd.InitDefaultHelpCmd()
	cmd.InitDefaultHelpFlag()

	buf := new(bytes.Buffer)
	name := cmd.CommandPath()

	if len(cmd.Long) > 0 {
		buf.WriteString("## Synopsis\n\n")
		buf.WriteString(cmd.Long + "\n\n")
	}

	if cmd.Runnable() {
		buf.WriteString(fmt.Sprintf("```\n%s\n```\n\n", cmd.UseLine()))
	}

	if len(cmd.Example) > 0 {
		buf.WriteString("## Examples\n\n")
		strings.Split(cmd.Example, "\n")

		blocks := strings.Split(cmd.Example, "\n\n\n")
		for _, block := range blocks {
			buf.WriteString(fmt.Sprintf("```bash\n%s\n```\n\n", block))
		}
	}

	if err := printOptions(buf, cmd, name); err != nil {
		return err
	}
	if hasSeeAlso(cmd) {
		buf.WriteString("## SEE ALSO\n\n")
		if cmd.HasParent() {
			parent := cmd.Parent()
			pname := parent.CommandPath()
			link := pname + markdownExtension
			link = strings.ReplaceAll(link, " ", "_")
			buf.WriteString(fmt.Sprintf("* [%s](%s)\t - %s\n", pname, linkHandler(link), parent.Short))
			cmd.VisitParents(func(c *cobra.Command) {
				if c.DisableAutoGenTag {
					cmd.DisableAutoGenTag = c.DisableAutoGenTag
				}
			})
		}

		children := cmd.Commands()
		sort.Sort(byName(children))

		for _, child := range children {
			if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
				continue
			}
			cname := name + " " + child.Name()
			link := cname + markdownExtension
			link = strings.ReplaceAll(link, " ", "_")
			buf.WriteString(fmt.Sprintf("* [%s](%s)\t - %s\n", cname, linkHandler(link), child.Short))
		}
		buf.WriteString("\n")
	}
	if !cmd.DisableAutoGenTag {
		buf.WriteString("###### Auto generated by spf13/cobra on " + time.Now().Format("2-Jan-2006") + "\n")
	}
	_, err := buf.WriteTo(w)
	return err
}

func printOptions(buf *bytes.Buffer, cmd *cobra.Command, name string) error {
	flags := cmd.NonInheritedFlags()
	flags.SetOutput(buf)
	if flags.HasAvailableFlags() {
		buf.WriteString("## Options\n\n```\n")
		flags.PrintDefaults()
		buf.WriteString("```\n\n")
	}

	parentFlags := cmd.InheritedFlags()
	parentFlags.SetOutput(buf)
	if parentFlags.HasAvailableFlags() {
		buf.WriteString("## Options inherited from parent commands\n\n```\n")
		parentFlags.PrintDefaults()
		buf.WriteString("```\n\n")
	}
	return nil
}

// Test to see if we have a reason to print See Also information in docs
// Basically this is a test for a parent command or a subcommand which is
// both not deprecated and not the autogenerated help command.
func hasSeeAlso(cmd *cobra.Command) bool {
	if cmd.HasParent() {
		return true
	}
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		return true
	}
	return false
}

type byName []*cobra.Command

func (s byName) Len() int           { return len(s) }
func (s byName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byName) Less(i, j int) bool { return s[i].Name() < s[j].Name() }
