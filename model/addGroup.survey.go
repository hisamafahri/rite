package model

import (
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
func AddGroupPrompt() []*survey.Question {
	return []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "What is your group name?"},
			Validate: survey.Required,
		},
		{
			Name: "path",
			Prompt: &survey.Input{
				Message: "Path to file?",
				Suggest: func(toComplete string) []string {
					files, _ := filepath.Glob(toComplete + "*")
					return files
				},
				Help: "A group cannot be empty and at least contain one file",
			},
		},
	}
}
