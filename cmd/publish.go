package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(publishCmd)
}

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publishes a cake",
	Long:  "Publishes a cake",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Enter the cake's path")
		}
		path_to_toml := args[0]

		if _, err := os.Stat(path_to_toml); os.IsNotExist(err) {
			return fmt.Errorf("That cake file does not exist")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}
