package prompt

import (
	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
func RemoveFilePrompt(fileList []string) []*survey.Question {
	return []*survey.Question{
		{
			Name: "file",
			Prompt: &survey.Select{
				Message: "Which file you want to remove?",
				Options: fileList,
			},
			Validate: survey.Required,
		},
		{
			Name: "confirm",
			Prompt: &survey.Confirm{
				Message: "Are you sure want to remove this file from those group(s)?",
			},
			Validate: survey.Required,
		},
	}
}
