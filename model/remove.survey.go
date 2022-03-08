package model

import (
	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
func RemovePrompt() []*survey.Question {
	return []*survey.Question{
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "What you want to remove?",
				Options: []string{"A Group", "A File", "A User"},
			},
			Validate: survey.Required,
		},
	}
}
