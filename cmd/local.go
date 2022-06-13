package cmd

import (
	"fmt"
	"os"

	"github.com/cake-cutter/cli/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(localCmd)
}

var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Creates a template using a local file",
	Long:  "Creates a template using a local file",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Enter the cake's path")
		}
		path_to_toml := args[0]

		if _, err := os.Stat(path_to_toml); os.IsNotExist(err) {
			return fmt.Errorf("File does not exist")
		}

		if len(args) < 2 {
			return fmt.Errorf("Enter the directory's path")
		}
		path_to_dir := args[1]
		if _, err := os.Stat(path_to_dir); os.IsNotExist(err) {
		} else {
			return fmt.Errorf("Directory already exist")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		path_to_toml := args[0]
		path_to_dir := args[1]

		var (
			conf *utils.Config
			err  error
		)

		fmt.Println(utils.Colorize("blue", "Creating template...\n "))

		utils.MakeItSpin(func() {
			conf, err = utils.ParseFromFile(path_to_toml)
			utils.Check(err)
		}, "Parsing toml...")

		utils.MakeItSpin(func() {
			err = utils.CutDir(path_to_dir, conf)
			utils.Check(err)
		}, "Cutting file structure...")

		utils.MakeItSpin(func() {
			err = utils.CutFiles(path_to_dir, conf)
			utils.Check(err)
		}, "Cutting files...")

		fmt.Println(utils.Colorize("green", "Successfully cutted `"+path_to_dir+"`"))

	},
}
