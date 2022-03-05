package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd will print the version of this CLI
var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the current version of the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Rite -- v0.1.0")
	},
}
