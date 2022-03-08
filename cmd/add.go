package cmd

import (
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hisamafahri/rite/helper"
	"github.com/hisamafahri/rite/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addUser command will add a user to specific group
var addUserCmd = &cobra.Command{
	Use:   "add",
	Short: "generate user and add the user to its desired groups",
	Run: func(cmd *cobra.Command, args []string) {

		userDetails := struct {
			FullName       string   `survey:"name"`
			Email          string   `survey:"email"`
			Password       string   `survey:"password"`
			SelectedGroups []string `survey:"group"`
			Expiration     string   `survey:"expire"`
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

		// perform the questions
		err = survey.Ask(model.AddUserPrompt(groupChoices), &userDetails)
		helper.CheckErr(err)

		expirationLength, err := strconv.Atoi(userDetails.Expiration)
		helper.CheckErr(err)

		// Set the expiration
		expiration := helper.Config{Expiry: time.Duration(expirationLength) * 24 * time.Hour}

		// create new key
		key, err := helper.CreateKey(userDetails.FullName, userDetails.Password, userDetails.Email, &expiration)
		helper.CheckErr(err)

		// armoring public key
		_, err = key.Armor()
		helper.CheckErr(err)

		// armoring private key
		_, err = key.ArmorPrivate(&expiration)
		helper.CheckErr(err)

		for _, selectedGroup := range userDetails.SelectedGroups {
			// Read the config file
			config, err := helper.LoadConfig()
			helper.CheckErr(err)

			// load users from config
			users := config.Users

			var tempMembersSlice []string
			var currentMembers interface{}
			var isGroupHasMembers bool

			// check if group listed in 'users:'
			if len(users) == 0 {
				users = make(map[string]interface{})
				currentMembers = users[selectedGroup]
				isGroupHasMembers = false
			} else {
				currentMembers, isGroupHasMembers = users[selectedGroup]
			}

			// if group have no member in the 'users' section
			// create an empty string, and append on it
			if !isGroupHasMembers {
				tempMembersSlice = append(tempMembersSlice, userDetails.Email)
				users[selectedGroup] = tempMembersSlice
			} else {
				var groupMembers interface{}
				groupMembers = currentMembers
				// convert groupMembers into []interface{}
				groupMembersSlice := currentMembers.([]interface{})

				// append new email to the groupMembersSlice
				groupMembers = append([]interface{}{userDetails.Email}, groupMembersSlice...)

				for groupName := range users {
					newMembers := groupMembers
					if groupName == selectedGroup {
						users[selectedGroup] = newMembers
					}
				}
			}

			// rewrite the config file
			viper.Set("users", users)

			err = viper.WriteConfig()
			helper.CheckErr(err)

			err = viper.WriteConfig()
			helper.CheckErr(err)

			// save the public and private key to a .pgp file
			ioutil.WriteFile(".rite/keys/"+strings.Replace(userDetails.Email, "@", ".", -1)+".public.gpg", key.Keyring(), 0666)
			ioutil.WriteFile(strings.Replace(userDetails.Email, "@", ".", -1)+".private.gpg", key.Secring(&expiration), 0666)
		}
	},
}
