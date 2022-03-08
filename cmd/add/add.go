package add

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/hisamafahri/rite/helper"
	prompt "github.com/hisamafahri/rite/prompt/add"
	"github.com/spf13/cobra"
)

// addUser command will add a user to specific group
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "generate user, create new group, add new files",
	Run: func(cmd *cobra.Command, args []string) {
		addType := struct {
			Type string `survey:"type"`
		}{}

		// perform the questions
		err := survey.Ask(prompt.AddPrompt(), &addType)
		helper.CheckErr(err)

		if addType.Type == "New Group" {
			addGroup()
		} else if addType.Type == "New Files" {
			addFile()
		} else if addType.Type == "New User" {
			addUser()
		}
	},
}
