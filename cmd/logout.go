package cmd

import (
	"github.com/cake-cutter/cli/cli/utils"
	"github.com/spf13/cobra"
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
