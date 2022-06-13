package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"will-change.later/utils"
)

func init() {
	rootCmd.AddCommand(userCmd)
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Checks if you are logged in with an account",
	Long:  "Checks if you are logged in with an account",

	Run: func(cmd *cobra.Command, args []string) {

		user, yes, err := utils.LoggedIn()
		utils.Check(err)

		if !yes {
			fmt.Println(utils.Colorize("red", "You are not logged in"))
			return
		}

		fmt.Println(utils.Colorize("green", "Logged in as "+*user))

	},
}
