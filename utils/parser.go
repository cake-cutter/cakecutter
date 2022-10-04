package utils

import (
	"encoding/json"
	"net/url"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Metadata struct {
		Name        string
		Description string
	} `toml:"metadata"`

	Content map[string]string `toml:"content"`

	FileStructure map[string]string `toml:"filestructure"`

	Commands map[int][2]string `toml:"toppings"`

	CommandsBefore map[int][2]string `toml:"batter"`

	Gatherers map[string]string `toml:"gatherers"`

	Questions map[string][]struct {
		Question string   `toml:"ques"`
		Type     string   `toml:"type"`
		Options  []string `toml:"options"`
		Default  string   `toml:"default"`
	} `toml:"questions"`
}

func ParseToml(txt string) (*Config, error) {
	var (
		conf = &struct {
			Metadata struct {
				Name        string
				Description string
			} `toml:"metadata"`

			Content map[string]string `toml:"content"`

			FileStructure map[string]string `toml:"filestructure"`

			Commands map[string][2]string `toml:"toppings"`

			CommandsBefore map[string][2]string `toml:"batter"`

			Gatherers map[string]string `toml:"gatherers"`

			Questions map[string][]struct {
				Question string   `toml:"ques"`
				Type     string   `toml:"type"`
				Options  []string `toml:"options"`
				Default  string   `toml:"default"`
			} `toml:"questions"`
		}{}
		err error
	)

	_, err = toml.Decode(txt, conf)
	if err != nil {
		return nil, err
	}

	_cp := make(map[int][2]string)
	for k, v := range conf.Commands {
		ik, err := strconv.Atoi(k)
		if err != nil {
			return nil, err
		}
		_cp[ik] = v
	}

	__cp := make(map[int][2]string)
	for k, v := range conf.CommandsBefore {
		ik, err := strconv.Atoi(k)
		if err != nil {
			return nil, err
		}
		__cp[ik] = v
	}

	cp := &Config{
		Metadata:       conf.Metadata,
		Content:        conf.Content,
		FileStructure:  conf.FileStructure,
		Commands:       _cp,
		CommandsBefore: __cp,
		Gatherers:      conf.Gatherers,
		Questions:      conf.Questions,
	}

	return cp, nil

}

func ParseFromFile(file string) (*Config, error) {

	f, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return ParseToml(string(f))
}

func ParseQuery(str string) (url.Values, error) {
	u, err := url.Parse("https://test/?" + str)
	if err != nil {
		return nil, err
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}

	return q, nil
}

func ParseUserJSON(body string) (*struct {
	Login string `json:"login"`
}, error) {
	var m struct {
		Login string `json:"login"`
	}
	err := json.Unmarshal([]byte(body), &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
