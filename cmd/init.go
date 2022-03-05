package cmd

import (
	"os"

	"github.com/hisamafahri/rite/data"
	"github.com/hisamafahri/rite/helper"
	"github.com/spf13/cobra"
)

// initCmd will initialize Rite in the repo
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Rite",
	Run: func(cmd *cobra.Command, args []string) {

		// Create the '.rite' folder
		err := os.Mkdir(".rite", 0755)
		helper.CheckErr(err)

		d1 := []byte(data.DefaultStarterFile)
		err = os.WriteFile(".rite/rite.config.yaml", d1, 0644)
		helper.CheckErr(err)

	},
}
