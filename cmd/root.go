package cmd

import (
	"os"

	"github.com/hisamafahri/rite/cmd/add"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rite",
	Short: "Rite enables you to store secrets in your Git repo, safely",
	Long: `Rite make it easy for you to encrypt specific files
with PGP encryption in a repo so they are "encrypted at 
rest" in your repository`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// add the 'version' or 'v' command
	rootCmd.AddCommand(versionCmd)

	// add the 'add' command
	rootCmd.AddCommand(add.AddCmd)

	// add the 'version' command
	rootCmd.AddCommand(initCmd)
}
