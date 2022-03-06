package cmd

import (
	"fmt"
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
		fmt.Println("Email Added: " + strings.Join(args, " "))

		viper.SetConfigName("rite.config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".rite")
		err := viper.ReadInConfig()
		if err != nil {
			helper.CheckErr(err)
		}

	},
}
