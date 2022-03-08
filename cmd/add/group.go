package add

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hisamafahri/rite/helper"
	"github.com/hisamafahri/rite/model"
)

func addGroup() {
	groupDetails := struct {
		Name string `survey:"name"`
		Path string `survey:"path"`
	}{}

	// perform the questions
	err := survey.Ask(model.AddGroupPrompt(), &groupDetails)
	helper.CheckErr(err)

	fmt.Println(groupDetails)

}
