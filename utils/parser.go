package utils

import (
	"encoding/json"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Metadata struct {
		Name        string
		Description string
	} `toml:"metadata"`

	Content map[string]string `toml:"content"`

	FileStructure map[string]string `toml:"filestructure"`

	Commands map[string]string `toml:"commands"`

	Questions map[string][]struct {
		Question string   `toml:"ques"`
		Type     string   `toml:"type"`
		Options  []string `toml:"options"`
		Default  string   `toml:"default"`
	} `toml:"questions"`
}

func ParseToml(txt string) (*Config, error) {
	var (
		conf = &Config{}
		err  error
	)

	_, err = toml.Decode(txt, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil

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
