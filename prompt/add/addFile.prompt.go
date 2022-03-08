package prompt

import (
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
func AddFilePrompt(groupChoices []string) []*survey.Question {
	return []*survey.Question{
		{
			Name: "path",
			Prompt: &survey.Input{
				Message: "Path to file?",
				Suggest: func(toComplete string) []string {
					files, _ := filepath.Glob(toComplete + "*")
					return files
				},
			},
		},
		{
			Name: "group",
			Prompt: &survey.MultiSelect{
				Message: "Which groups this file should belong to?",
				Options: groupChoices,
			},
			Validate: survey.Required,
		},
	}
}
