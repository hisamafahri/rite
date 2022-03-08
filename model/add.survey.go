package model

import (
	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
func AddPrompt() []*survey.Question {
	return []*survey.Question{
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "What you want to add?",
				Options: []string{"New Group", "New Files", "New User"},
			},
			Validate: survey.Required,
		},
	}
}
