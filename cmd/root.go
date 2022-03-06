package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// group flag variable for 'add-user' command
var Group string

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

	// add the 'generate' or 'g' command
	rootCmd.AddCommand(generateCmd)

	// add the 'add-user' command
	// and '--group' flag
	rootCmd.AddCommand(addUserCmd)
	addUserCmd.PersistentFlags().StringVarP(&Group, "group", "g", "", "Group the new user should be in. Separated with comma (,).")
	addUserCmd.MarkPersistentFlagRequired("group")

	// add the 'version' command
	rootCmd.AddCommand(initCmd)
}
