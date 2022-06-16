package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
  "strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cake-cutter/cc/utils"
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

		user, loggedIn, err := utils.LoggedIn()
		utils.Check(err)

		if !loggedIn {
			fmt.Println(utils.Colorize("red", "You are not logged in"))
			os.Exit(1)
		}

		fmt.Println(utils.Colorize("blue", "Publishing cake..."))

		path_to_toml := args[0]

		cake, err := utils.ParseFromFile(path_to_toml)
		utils.Check(err)

		fmt.Println("\nEnsure cake details...")

		fmt.Println(utils.Colorize("blue", " - Name: ") + utils.Colorize("gray", cake.Metadata.Name))
		fmt.Println(utils.Colorize("blue", " - Description: ") + utils.Colorize("gray", cake.Metadata.Description))
		fmt.Println()

		var result string

		err = survey.AskOne(
			&survey.Select{
				Message: "Is everything correct?",
				Options: []string{"Yes", "No"},
			},
			&result,
		)
		utils.Check(err)

		if result == "No" {
			fmt.Println(utils.Colorize("red", "\nAborted"))
			os.Exit(1)
		}

    if strings.Contains(cake.Metadata.Name, " ") {
      fmt.Println(utils.Colorize("red", "\nThe name cannot contain whitespace!"))
      os.Exit(1)
    }

    if len(cake.Metadata.Name) < 1 {
      fmt.Println(utils.Colorize("red", "\nThe name is empty"))
      os.Exit(1)
    }

		utils.Input(
			"Enter the REAMDE's path ",
			"README.md",
			&result,
			func(s string) error {
				exists, err := utils.PathExists(s)
				if err != nil {
					return err
				}
				if !exists {
					return fmt.Errorf("That path does not exist")
				}
				return nil
			},
		)

		fmt.Println()

		utils.MakeItSpin(func() {

			content, err := os.ReadFile(path_to_toml)
			utils.Check(err)

			body, err := os.ReadFile(result)
			utils.Check(err)

			j_ := map[string]string{
				"name":   cake.Metadata.Name,
				"short":  cake.Metadata.Description,
				"cake":   string(content),
				"dsc":    string(body),
				"author": *user,
			}

			aa, err := json.Marshal(j_)
			utils.Check(err)

			resp, err := http.Post(utils.BackendURL+"/publish", "application/json", bytes.NewBuffer(aa))
			utils.Check(err)

			defer resp.Body.Close()

			b, err := io.ReadAll(resp.Body)
			utils.Check(err)

			var j struct {
				Success int    `json:"success"`
				Data    string `json:"data"`
				Error   string `json:"error"`
			}

			err = json.Unmarshal([]byte(b), &j)
			utils.Check(err)

			fmt.Println()

			if resp.StatusCode != 200 {
				if j.Success == 1 {
					fmt.Println(utils.Colorize("red", j.Error))
					os.Exit(1)
				}
			} else {
				if j.Success == 0 {
					fmt.Println(utils.Colorize("green", "Cake published successfully"))
					os.Exit(0)
				}
			}

		}, "Publishing...")

	},
}
