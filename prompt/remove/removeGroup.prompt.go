package prompt

import (
	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
func RemoveGroupPrompt(groupChoices []string) []*survey.Question {
	return []*survey.Question{
		{
			Name: "group",
			Prompt: &survey.MultiSelect{
				Message: "Which groups you want to delete?",
				Options: groupChoices,
			},
			Validate: survey.Required,
		},
		{
			Name: "confirm",
			Prompt: &survey.Confirm{
				Message: "Are you sure want to delete those group(s) and its member files?",
			},
			Validate: survey.Required,
		},
	}
}
