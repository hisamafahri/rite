package add

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hisamafahri/rite/helper"
	"github.com/hisamafahri/rite/model"
	"github.com/spf13/viper"
)

func addFile() {
	newFileDetails := struct {
		Path           string   `survey:"path"`
		SelectedGroups []string `survey:"group"`
	}{}

	config, err := helper.LoadConfig()
	helper.CheckErr(err)

	// load groups from config
	existGroups := config.Groups

	// get the group names in the config file
	var groupChoices []string
	for groupName := range existGroups {
		groupChoices = append(groupChoices, groupName)
	}

	if len(groupChoices) <= 0 {
		helper.CheckErr(errors.New("rite: no group found in the config file. Please create a group first"))
		return
	}

	// perform the questions
	err = survey.Ask(model.AddFilePrompt(groupChoices), &newFileDetails)
	helper.CheckErr(err)

	for _, group := range newFileDetails.SelectedGroups {
		fileMembers := existGroups[group]
		fileMembersSlice := fileMembers.([]interface{})

		// append new email to the groupMembersSlice
		fileMembers = append([]interface{}{newFileDetails.Path}, fileMembersSlice...)
		existGroups[group] = fileMembers
	}

	viper.Set("groups", existGroups)

	err = viper.WriteConfig()
	helper.CheckErr(err)

	err = viper.WriteConfig()
	helper.CheckErr(err)
}
