package remove

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hisamafahri/rite/helper"
	prompt "github.com/hisamafahri/rite/prompt/remove"
	"github.com/spf13/viper"
)

func removeGroup() {
	groupDetails := struct {
		Groups  []string `survey:"group"`
		Confirm bool     `survey:"confirm"`
	}{}

	config, err := helper.LoadConfig()
	helper.CheckErr(err)

	// load groups from config
	groupsList := config.Groups

	// load users list
	usersList := config.Users

	// get the group names in the config file
	var groupChoices []string
	for groupName := range groupsList {
		groupChoices = append(groupChoices, groupName)
	}

	// return error if no group in the config file
	if len(groupChoices) <= 0 {
		helper.CheckErr(errors.New("rite: no group found in the config file. Please create a group first"))
		return
	}

	// perform the questions
	err = survey.Ask(prompt.RemoveGroupPrompt(groupChoices), &groupDetails)
	helper.CheckErr(err)

	// close the prompt if the confirmation return is N or no
	if !groupDetails.Confirm {
		helper.CheckErr(errors.New("rite: process cancelled"))
		return
	}

	// loop through the selected groups and delete the files
	for _, selectedGroup := range groupDetails.Groups {
		_, isGroupExist := groupsList[selectedGroup]
		_, isUsersExist := usersList[selectedGroup]

		// making sure the group is exist
		if isGroupExist {
			delete(viper.Get("groups").(map[string]interface{}), selectedGroup)
			if isUsersExist {
				delete(viper.Get("users").(map[string]interface{}), selectedGroup)
			}
		}
	}

	// rewrite the config file
	err = viper.WriteConfig()
	helper.CheckErr(err)
}
