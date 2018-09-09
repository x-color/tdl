package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var cfgFile string
var todosFile = ""

func init() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	todosFile = home + "/.tdl/.todos.json"
}

func newCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tdl",
		Short: "This tool is todos manager.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(newCmdAdd())
	cmd.AddCommand(newCmdDelete())
	cmd.AddCommand(newCmdCheck())
	cmd.AddCommand(newCmdStar())
	cmd.AddCommand(newCmdEdit())
	cmd.AddCommand(newCmdShow())

	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cmd := newCmdRoot()
	cmd.SetOutput(os.Stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}
