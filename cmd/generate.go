package cmd

import (
	"io/ioutil"
	"time"

	"github.com/hisamafahri/rite/helper"
	"github.com/spf13/cobra"
)

// generateCmd willl generate a new public and private gpg for a new user
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "generate a new public and private gpg for a new user",
	Run: func(cmd *cobra.Command, args []string) {

		// Set the expiration
		config := helper.Config{Expiry: 365 * 24 * time.Hour}
		key, err := helper.CreateKey("Joe Smith", "test key", "joe@example.com", &config)
		helper.CheckErr(err)

		_, err = key.Armor()
		helper.CheckErr(err)
		// fmt.Printf("%s\n", pub)

		_, err = key.ArmorPrivate(&config)
		helper.CheckErr(err)
		// fmt.Printf("%s\n", priv)

		ioutil.WriteFile(".rite/keys/joe.example.com.public.gpg", key.Keyring(), 0666)
		ioutil.WriteFile("joe.example.com.private.gpg", key.Secring(&config), 0666)
	},
}
