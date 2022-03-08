package add

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hisamafahri/rite/helper"
	"github.com/hisamafahri/rite/model"
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
		err := survey.Ask(model.AddPrompt(), &addType)
		helper.CheckErr(err)

		if addType.Type == "New Groups" {
			fmt.Print("New Groups")
		} else if addType.Type == "New Files" {
			fmt.Println("new Files")
		} else {
			addUser()
		}
	},
}
