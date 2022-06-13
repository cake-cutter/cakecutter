package cmd

import (
	"github.com/spf13/cobra"
	"will-change.later/utils"
)

func init() {
	rootCmd.AddCommand(logoutCmd)
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logs you out",
	Long:  "Logs you out",

	Run: func(cmd *cobra.Command, args []string) {

		err := utils.Logout()
		utils.Check(err)

	},
}
