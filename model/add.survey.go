package model

import (
	"errors"
	"regexp"
	"unicode"

	"github.com/AlecAivazis/survey/v2"
)

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func isNumber(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// the questions to ask
func AddUserPrompt(groupChoices []string) []*survey.Question {
	return []*survey.Question{
		{
			Name:      "name",
			Prompt:    &survey.Input{Message: "What is your full name?"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name:   "email",
			Prompt: &survey.Input{Message: "What is your email address?"},
			Validate: func(val interface{}) error {
				if str, ok := val.(string); !ok || !isEmailValid(str) {
					return errors.New("input is not valid email address")
				}
				return nil
			},
		},
		{
			Name: "password",
			Prompt: &survey.Password{
				Message: "Your password?",
			},
			Validate: survey.MinLength(8),
		},
		{
			Name: "group",
			Prompt: &survey.MultiSelect{
				Message: "Which groups this user should belong to?",
				Options: groupChoices,
			},
			Validate: survey.Required,
		},
		{
			Name: "expire",
			Prompt: &survey.Input{
				Message: "Expires in (days)?",
				Default: "365",
			},
			Validate: func(val interface{}) error {
				if str, ok := val.(string); !ok || !isNumber(str) {
					return errors.New("input is not valid number")
				}
				return nil
			},
		},
	}
}
