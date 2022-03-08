package add

import (
	"github.com/spf13/cobra"
)

// addUser command will add a user to specific group
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "generate user, create new group, add new files",
	Run: func(cmd *cobra.Command, args []string) {
		addUser()
	},
}
