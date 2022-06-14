package utils

import (
	"os"
	"os/exec"
	"strings"
)

func RemoveItemFromSlice(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func CutDir(dir string, conf *Config) error {

	for k, v := range conf.FileStructure {

		if k == "root" {
			k = ""
		}
		ke := strings.ReplaceAll(strings.ReplaceAll(k, "--", "/"), "-", ".")

		err := os.MkdirAll(dir+"/"+ke, 0755)
		if err != nil {
			return err
		}

		for _, va := range v {
			_, err = os.Create(dir + "/" + ke + "/" + va)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func CutFiles(dir string, conf *Config) error {

	for k, v := range conf.Content {

		ke := strings.ReplaceAll(strings.ReplaceAll(k, "--", "/"), "-", ".")

		file, err := os.Create(dir + "/" + ke)
		if err != nil {
			return err
		}

		if _, err = file.WriteString(v); err != nil {
			return err
		}

	}

	return nil

}

func CutDaCommands(dir string, cmds map[string][]string) error {

	for _, cmd := range cmds {
		cmd := exec.Command(cmd[0], cmd[1:]...)
		cmd.Dir = dir
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil

}
