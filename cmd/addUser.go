package cmd

import (
	"errors"
	"strings"

	"github.com/hisamafahri/rite/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addUser command will add a user to specific group
var addUserCmd = &cobra.Command{
	Use:   "add-user [email]",
	Short: "add a user to specific group",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// get the new email value
		newEmail := strings.Join(args, " ")

		// get the '--group' flag value
		groupFlag, _ := cmd.Flags().GetString("group")

		// remove all spaces
		groupFlag = strings.ReplaceAll(groupFlag, " ", "")

		// split the groupFlag into slice
		designatedGroups := strings.Split(groupFlag, ",")

		for _, designatedGroup := range designatedGroups {
			// Read the config file
			config, err := helper.LoadConfig()
			helper.CheckErr(err)

			// load users from config
			users := config.Users

			// check if group exist in config file
			groupMembers, isGroupExist := users[designatedGroup]

			// return error if group doesn't exist
			if !isGroupExist {
				helper.CheckErr(errors.New("rite: group " + designatedGroup + " doesn't exist in the .rite/config.rite.yaml file"))
				return
			}

			// convert groupMembers into []interface{}
			groupMembersSlice := groupMembers.([]interface{})

			// append new email to the groupMembersSlice
			groupMembers = append([]interface{}{newEmail}, groupMembersSlice...)

			for groupName := range users {
				// fmt.Println("Received ID:", v, "index: ", i)
				newMembers := groupMembers
				if groupName == designatedGroup {
					users[designatedGroup] = newMembers
				}
			}

			// rewrite the config file
			viper.Set("users", users)

			err = viper.WriteConfig()
			helper.CheckErr(err)

			err = viper.WriteConfig()
			helper.CheckErr(err)
		}
	},
}
