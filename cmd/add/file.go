package add

import (
	"errors"
	"reflect"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hisamafahri/rite/helper"
	prompt "github.com/hisamafahri/rite/prompt/add"
	"github.com/spf13/viper"
)

func uniqueList(src interface{}) interface{} {
	srcv := reflect.ValueOf(src)
	dstv := reflect.MakeSlice(srcv.Type(), 0, 0)
	visited := make(map[interface{}]struct{})
	for i := 0; i < srcv.Len(); i++ {
		elemv := srcv.Index(i)
		if _, ok := visited[elemv.Interface()]; ok {
			continue
		}
		visited[elemv.Interface()] = struct{}{}
		dstv = reflect.Append(dstv, elemv)
	}
	return dstv.Interface()
}

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
	err = survey.Ask(prompt.AddFilePrompt(groupChoices), &newFileDetails)
	helper.CheckErr(err)

	for _, group := range newFileDetails.SelectedGroups {
		fileMembers := existGroups[group]
		fileMembersSlice := fileMembers.([]interface{})

		for _, file := range fileMembersSlice {
			if file == newFileDetails.Path {
				helper.CheckErr(errors.New("rite: file '" + newFileDetails.Path + "' is skipped because it's already member of '" + group + "' group"))
			}
		}

		// append new email to the groupMembersSlice
		fileMembers = append([]interface{}{newFileDetails.Path}, fileMembersSlice...)
		uniqueList := uniqueList(fileMembers)
		existGroups[group] = uniqueList
	}

	viper.Set("groups", existGroups)

	err = viper.WriteConfig()
	helper.CheckErr(err)

	err = viper.WriteConfig()
	helper.CheckErr(err)
}
