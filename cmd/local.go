package cmd

import (
	"fmt"
	"os"
	"runtime"
	"sort"

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
		gatherers := make(map[string]string)
		theData := utils.Data{
			Os: runtime.GOOS,
		}

		err = utils.CutTheGatherers(gatherers, conf)
		utils.Check(err)

		err = utils.CutTheQuestions(&ans, conf)
		utils.Check(err)

		theData.Ans = ans
		theData.Gatherer = gatherers

		_cs, err := utils.ParseCommands(conf.CommandsBefore, theData)
		utils.Check(err)

		sort.Ints(_cs)

		if len(_cs) > 0 {
			fmt.Println("\n" + utils.Colorize("green", "These commands are going to run... If these commands seems suspicious or harmful please report them by making an issue on the repo - `https://github.com/cake-cutter/cakes.run`"))
			for _, v := range _cs {
				fmt.Println(utils.Colorize("gray", "  "+conf.CommandsBefore[v][0]))
			}

			var _res string

			survey.AskOne(&survey.Select{
				Message: "Do you want to run these commands?",
				Options: []string{"Yes", "No"},
			}, &_res)

			if _res == "No" {
				fmt.Println("\n" + utils.Colorize("red", "Aborted!"))
				os.Exit(0)
			}

			fmt.Println()

		}

		path_exists, err := utils.PathExists(path_to_dir)
		utils.Check(err)
		if !path_exists {
			err = os.Mkdir(path_to_dir, 0755)
			utils.Check(err)
		}

		utils.MakeItSpin(func() {
			err = utils.CutDaCommands(path_to_dir, conf.CommandsBefore, _cs)
			utils.Check(err)
		}, "Folding the batter...")

		utils.MakeItSpin(func() {
			err = utils.CutDir(path_to_dir, conf, theData)
			utils.Check(err)
		}, "Cutting file structure...")

		utils.MakeItSpin(func() {
			err = utils.CutFiles(path_to_dir, conf, theData)
			utils.Check(err)
		}, "Cutting files...")

		cs, err := utils.ParseCommands(conf.Commands, theData)
		utils.Check(err)

		sort.Ints(cs)

		if len(cs) > 0 {
			fmt.Println("\n" + utils.Colorize("green", "These commands are gonna run to finish the setup... If these commands seems suspicious or harmful please report them by making an issue on the repo - `https://github.com/cake-cutter/cakes.run`"))
			for _, v := range cs {
				fmt.Println(utils.Colorize("gray", "  "+conf.Commands[v][0]))
			}

			var _res string

			survey.AskOne(&survey.Select{
				Message: "Do you want to run these commands?",
				Options: []string{"Yes", "No"},
			}, &_res)

			if _res == "No" {
				fmt.Println("\n" + utils.Colorize("red", "Aborted!"))
				fmt.Println(utils.Colorize("green", "The rest of the cake has been cut."))
				os.Exit(0)
			}

			fmt.Println()
		}

		utils.MakeItSpin(func() {
			err = utils.CutDaCommands(path_to_dir, conf.Commands, cs)
			utils.Check(err)
		}, "Sprinkling toppings...")

		fmt.Println(utils.Colorize("green", "Successfully cut `"+path_to_dir+"`"))

	},
}
