package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cake-cutter/cc/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(localCmd)
}

var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Creates a template using a local file",
	Long:  "Creates a template using a local file. This is useful for creating templates from a TOML Cakefile file you created locally. Read the docs at https://docs.cakes.run/2-usage/ to see how to create a Cakefile",

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

			dir, err := os.ReadDir(path_to_dir)
			utils.Check(err)
			if len(dir) > 0 {

				fmt.Println()

				var res string

				survey.AskOne(&survey.Select{
					Message: "The directory `" + path_to_dir + "` has some files. Are you sure you want to use it?",
					Options: []string{"Yes", "No"},
					Default: "No",
				}, &res)

				if res == "No" {
					return fmt.Errorf("Aborted")
				}

				fmt.Println()

				return nil

			}
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

		ans := make(map[string]string)

		err = utils.CutTheQuestions(&ans, conf)
		utils.Check(err)

		utils.MakeItSpin(func() {
			err = utils.CutDir(path_to_dir, conf, ans)
			utils.Check(err)
		}, "Cutting file structure...")

		utils.MakeItSpin(func() {
			err = utils.CutFiles(path_to_dir, conf, ans)
			utils.Check(err)
		}, "Cutting files...")

		utils.MakeItSpin(func() {
			err = utils.CutDaCommands(path_to_dir, conf.Commands, ans)
			utils.Check(err)
		}, "Cutting commands...")

		fmt.Println(utils.Colorize("green", "Successfully cut `"+path_to_dir+"`"))

	},
}
