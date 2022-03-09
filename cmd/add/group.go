package add

import (
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hisamafahri/rite/helper"
	prompt "github.com/hisamafahri/rite/prompt/add"
	"github.com/spf13/viper"
)

func addGroup() {
	groupDetails := struct {
		Name string `survey:"name"`
		Path string `survey:"path"`
	}{}

	// perform the questions
	err := survey.Ask(prompt.AddGroupPrompt(), &groupDetails)
	helper.CheckErr(err)

	// Read the config file
	config, err := helper.LoadConfig()
	helper.CheckErr(err)

	// load users from config
	groups := config.Groups

	var isGroupNameExist bool

	if len(groups) == 0 {
		groups = make(map[string]interface{})
	} else {
		_, isGroupNameExist = groups[groupDetails.Name]
	}

	if isGroupNameExist {
		fmt.Println(errors.New("rite: group '" + groupDetails.Name + "' already exist"))
		return
	}

	// add new group
	groups[groupDetails.Name] = []string{groupDetails.Path}

	// rewrite the config file
	viper.Set("groups", groups)

	err = viper.WriteConfig()
	helper.CheckErr(err)
}
