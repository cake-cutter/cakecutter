package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/template"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mattn/go-shellwords"
)

type Data struct {
	Ans      map[string]string
	Gatherer map[string]string
	Os       string
}

func CutTheQuestions(ans *map[string]string, conf *Config) error {

	answ := make(map[string]string)

	for k, v := range conf.Questions {

		switch v[0].Type {

		case "input":

			fmt.Print(Colorize("green", "? ") + v[0].Question + " ")
			if v[0].Default != "" {
				fmt.Print(Colorize("gray", "("+v[0].Default+") ") + colorBlue)
			}

			scanner := bufio.NewScanner(os.Stdin)
			var result string

			if scanner.Scan() {
				if scanner.Text() != "" {
					result = scanner.Text()
				} else {
					result = v[0].Default
				}
			}

			fmt.Print(colorReset)

			answ[k] = result

		case "menu", "select":

			var result string

			err := survey.AskOne(
				&survey.Select{
					Message: v[0].Question,
					Options: v[0].Options,
				},
				&result,
			)

			if err != nil {
				return err
			}

			answ[k] = result

		}
	}

	*ans = answ

	return nil
}

func RemoveItemFromSlice(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func CutDir(dir string, conf *Config, ans Data) error {

	var err error

	path_exists, err := PathExists(dir)
	if err != nil {
		return err
	}

	if !path_exists {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}

	for k, v := range conf.FileStructure {

		buf := &bytes.Buffer{}

		tm := template.New(k)
		tm, err = tm.Parse(v)
		if err != nil {
			return err
		}

		err = tm.Execute(buf, ans)
		if err != nil {
			return err
		}

		r := buf.String()

		if strings.Contains(fmt.Sprint(r), "true") {
			if strings.HasSuffix(k, "/") {
				err := os.MkdirAll(dir+"/"+k, 0755)
				if err != nil {
					return err
				}
			} else {
				if strings.Contains(k, "/") {
					err = os.MkdirAll(dir+"/"+strings.Join(RemoveItemFromSlice(strings.Split(k, "/"), len(strings.Split(k, "/"))-1), "/"), 0755)
					if err != nil {
						return err
					}
				}
				_, err = os.Create(dir + "/" + k)

				if err != nil {
					return err
				}
			}
		}

	}

	return nil
}

func CutFiles(dir string, conf *Config, ans Data) error {

	for k, v := range conf.Content {

		exists, err := PathExists(dir + "/" + k)
		if err != nil {
			return err
		}

		if exists {
			buf := &bytes.Buffer{}

			file, err := os.Create(dir + "/" + k)
			if err != nil {
				return err
			}

			tm := template.New(k)
			tm, err = tm.Parse(v)
			if err != nil {
				return err
			}

			err = tm.Execute(buf, ans)
			if err != nil {
				return err
			}

			r := buf.String()

			if _, err = file.WriteString(r); err != nil {
				return err
			}
		}

	}

	return nil

}

func ParseCommands(cmds map[int][2]string, ans Data) ([]int, error) {

	var res []int

	for k, v := range cmds {

		buf := &bytes.Buffer{}

		tm := template.New(v[0])
		tm, err := tm.Parse(v[1])

		if err != nil {
			return nil, err
		}

		err = tm.Execute(buf, ans)
		if err != nil {
			return nil, err
		}

		r := buf.String()

		if strings.Contains(r, "true") {
			res = append(res, k)
		}

	}

	sort.Ints(res)

	return res, nil

}

func CutDaCommands(dir string, cmds map[int][2]string, order []int) error {

	sort.Ints(order)

	for _, v := range order {

		cmands, err := shellwords.Parse(cmds[v][0])
		if err != nil {
			return err
		}
		cmd := exec.Command(cmands[0], cmands[1:]...)
		cmd.Dir = dir
		err = cmd.Run()
		if err != nil {
			return err
		}

	}

	return nil

}

func Input(ques string, def string, ans *string, validate func(string) error) error {

	fmt.Print(Colorize("green", "? ") + "Path to REAMDE.md?" + " ")
	if def != "" {
		fmt.Print(Colorize("gray", "("+def+") ") + colorBlue)
	}

	var result string

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		if scanner.Text() != "" {
			result = scanner.Text()
		} else {
			result = "README.md"
		}
	}

	err := validate(result)
	if err != nil {
		fmt.Println(Colorize("red", "✖ ") + err.Error())
		Input(ques, def, ans, validate)
		return nil
	}

	*ans = result

	return nil

}
