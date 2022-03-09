package remove

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hisamafahri/rite/helper"
	prompt "github.com/hisamafahri/rite/prompt/remove"
)

func removeFile() {
	fileDetails := struct {
		File    string   `survey:"file"`
		Group   []string `survey:"group"`
		Confirm bool     `survey:"confirm"`
	}{}

	exampleFiles := []string{"file 1", "file 2", "file 3"}
	exampleGroups := []string{"group 1", "group 2", "group 3"}

	// perform the questions
	err := survey.Ask(prompt.RemoveFilePrompt(exampleFiles, exampleGroups), &fileDetails)
	helper.CheckErr(err)
	fmt.Println(fileDetails)
}
