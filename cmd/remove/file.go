package remove

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hisamafahri/rite/helper"
	prompt "github.com/hisamafahri/rite/prompt/remove"
)

func removeFile() {
	fileDetails := struct {
		File    string `survey:"file"`
		Confirm bool   `survey:"confirm"`
	}{}

	exampleFiles := []string{"file 1", "file 2", "file 3"}

	// perform the questions
	err := survey.Ask(prompt.RemoveFilePrompt(exampleFiles), &fileDetails)
	helper.CheckErr(err)
	fmt.Println(fileDetails)
}
