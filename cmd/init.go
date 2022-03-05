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

		// Create the 'keys' folder inside '.rite' directory
		err = os.Mkdir(".rite/keys", 0755)
		helper.CheckErr(err)

		// Create the 'rite.config.yaml' file inside the '.rite' folder
		d1 := []byte(data.DefaultStarterFile)
		err = os.WriteFile(".rite/rite.config.yaml", d1, 0644)
		helper.CheckErr(err)

		// If the file doesn't exist, create it, or append to the file
		f, err := os.OpenFile(".gitignore",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			helper.CheckErr(err)
		}
		defer f.Close()
		if _, err := f.WriteString("\n\n# Ignore the private gpg file\n*.private.gpg"); err != nil {
			helper.CheckErr(err)
		}

	},
}
