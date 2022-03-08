package remove

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hisamafahri/rite/helper"
	"github.com/hisamafahri/rite/model"
	"github.com/spf13/cobra"
)

// addUser command will add a user to specific group
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a user, group, or files",
	Run: func(cmd *cobra.Command, args []string) {
		removeType := struct {
			Type string `survey:"type"`
		}{}

		// perform the questions
		err := survey.Ask(model.RemovePrompt(), &removeType)
		helper.CheckErr(err)

		if removeType.Type == "A Group" {
			fmt.Println(removeType.Type)
		} else if removeType.Type == "A File" {
			fmt.Println(removeType.Type)
		} else if removeType.Type == "A User" {
			fmt.Println(removeType.Type)
		}
	},
}
