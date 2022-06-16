package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cake-cutter/cc/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cutCmd)
}

var cutCmd = &cobra.Command{
	Use:   "cut",
	Short: "Cuts a online cake xD (Creates a template from an online template)",
	Long:  "Cuts a online cake xD (Creates a template from an online template)",
  Aliases: []string{"run"},

	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) < 1 {
			return fmt.Errorf("Enter the cake's name!")
		}
		if len(args) < 2 {
			return fmt.Errorf("Enter the directory's path!")
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

		path_to_dir := args[1]

		var (
			conf *utils.Config
			err  error
		)

		fmt.Println(utils.Colorize("blue", "\nCutting cake...\n "))

		utils.MakeItSpin(func() {

			resp, err := http.Get(utils.BackendURL + "/get?name=" + args[0])
			utils.Check(err)

			var j struct {
				Success int    `json:"success"`
				Error   string `json:"error"`
				Data    struct {
					Name   string `json:"name"`
					Short  string `json:"short"`
					Dsc    string `json:"dsc"`
					Author string `json:"author"`
					Cake   string `json:"cake"`
				} `json:"data"`
			}

			defer resp.Body.Close()

			b, err := io.ReadAll(resp.Body)
			utils.Check(err)

			err = json.Unmarshal([]byte(b), &j)
			utils.Check(err)

			if resp.StatusCode != 200 {
				if j.Success == 1 {
					fmt.Println(utils.Colorize("red", j.Error))
					os.Exit(1)
				}
			}

			conf, err = utils.ParseToml(j.Data.Cake)
			utils.Check(err)

		}, "Fetching cake...")

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
