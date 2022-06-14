package utils

import (
	"os"
	"os/exec"
)

func CutTheQuestions(ans *map[string]string, conf *Config) error {

	return nil
}

func RemoveItemFromSlice(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func CutDir(dir string, conf *Config) error {

	for k, v := range conf.FileStructure {

		if k == "root" {
			k = ""
		}

		err := os.MkdirAll(dir+"/"+k, 0755)
		if err != nil {
			return err
		}

		for _, va := range v {
			_, err = os.Create(dir + "/" + k + "/" + va)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func CutFiles(dir string, conf *Config) error {

	for k, v := range conf.Content {

		file, err := os.Create(dir + "/" + k)
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
