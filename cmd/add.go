package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/hisamafahri/rite/helper"
	"github.com/hisamafahri/rite/model"
	"github.com/spf13/cobra"
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

		fmt.Printf("name: %s\n", userDetails.FullName)
		fmt.Printf("email: %s\n", userDetails.Email)
		fmt.Printf("password: %s\n", userDetails.Password)
		fmt.Printf("groups: %s\n", userDetails.SelectedGroups)
		fmt.Printf("expire: %s\n", userDetails.Expiration)

		// // Set the expiration
		// config := helper.Config{Expiry: 365 * 24 * time.Hour}

		// // create new key
		// key, err := helper.CreateKey("Joe Smith", "test key", "joe@example.com", &config)
		// helper.CheckErr(err)

		// // armoring public key
		// _, err = key.Armor()
		// helper.CheckErr(err)

		// // armoring private key
		// _, err = key.ArmorPrivate(&config)
		// helper.CheckErr(err)

		// // save the public and private key to a .pgp file
		// ioutil.WriteFile(".rite/keys/joe.example.com.public.gpg", key.Keyring(), 0666)
		// ioutil.WriteFile("joe.example.com.private.gpg", key.Secring(&config), 0666)

		// // get the new email value
		// newEmail := strings.Join(args, " ")

		// // get the '--group' flag value
		// groupFlag, _ := cmd.Flags().GetString("group")

		// // remove all spaces
		// groupFlag = strings.ReplaceAll(groupFlag, " ", "")

		// // split the groupFlag into slice
		// designatedGroups := strings.Split(groupFlag, ",")

		// for _, designatedGroup := range designatedGroups {
		// // Read the config file
		// config, err := helper.LoadConfig()
		// helper.CheckErr(err)

		// // load users from config
		// users := config.Users

		// 	// check if group exist in config file
		// 	groupMembers, isGroupExist := users[designatedGroup]

		// 	// return error if group doesn't exist
		// 	if !isGroupExist {
		// 		helper.CheckErr(errors.New("rite: group " + designatedGroup + " doesn't exist in the .rite/config.rite.yaml file"))
		// 		return
		// 	}

		// 	// convert groupMembers into []interface{}
		// 	groupMembersSlice := groupMembers.([]interface{})

		// 	// append new email to the groupMembersSlice
		// 	groupMembers = append([]interface{}{newEmail}, groupMembersSlice...)

		// 	for groupName := range users {
		// 		// fmt.Println("Received ID:", v, "index: ", i)
		// 		newMembers := groupMembers
		// 		if groupName == designatedGroup {
		// 			users[designatedGroup] = newMembers
		// 		}
		// 	}

		// 	// rewrite the config file
		// 	viper.Set("users", users)

		// 	err = viper.WriteConfig()
		// 	helper.CheckErr(err)

		// 	err = viper.WriteConfig()
		// 	helper.CheckErr(err)
		// }
	},
}
