package cmd

import (
	"fmt"

	"github.com/cake-cutter/cc/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your account",
	Long:  "Login to your account your github",

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Logging in with github...\n", utils.Colorize("blue", "\nPress `Enter` to login"))

		fmt.Scanln()

		utils.Login()

	},
}
