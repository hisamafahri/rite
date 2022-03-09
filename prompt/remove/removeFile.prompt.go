package prompt

import (
	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
func RemoveFilePrompt(fileChoices []string, groupChoices []string) []*survey.Question {
	return []*survey.Question{
		{
			Name: "file",
			Prompt: &survey.Select{
				Message: "Which file you want to remove?",
				Options: fileChoices,
			},
			Validate: survey.Required,
		},
		{
			Name: "group",
			Prompt: &survey.MultiSelect{
				Message: "From which groups you want to delet that file?",
				Options: groupChoices,
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
