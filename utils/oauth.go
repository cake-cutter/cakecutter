package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/pkg/browser"
)

var BackendURL = "https://cakes.run/api"

func Login() {
	var (
		body []byte
		q    url.Values
	)

	resp, err := http.PostForm("https://github.com/login/device/code", url.Values{
		"client_id": {"8a0e6a8495e7ce5c4218"},
	})
	Check(err)

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	Check(err)

	q, err = ParseQuery(string(body))
	Check(err)

	fmt.Printf("Opening %s \nEnter this code - %s\n", Colorize("green", "https://github.com/login/device"), Colorize("blue", q.Get("user_code")))

	browser.OpenURL("https://github.com/login/device")

	fmt.Println("\nType enter to continue -")
	fmt.Scanln()

	var res *http.Response

	res, err = http.PostForm("https://github.com/login/oauth/access_token", url.Values{
		"client_id":   {"8a0e6a8495e7ce5c4218"},
		"device_code": {q.Get("device_code")},
		"grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
	})

	if err != nil {
		ClearScreen()
		fmt.Println("\nError: " + Colorize("red", err.Error()))
		Login()
		return
	}

	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	Check(err)

	var qa url.Values

	qa, err = ParseQuery(string(body))
	Check(err)

	if qa.Get("error") != "" {
		fmt.Println("\nError: " + Colorize("red", qa.Get("error")))
		fmt.Println()
		Login()
		return
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Set("Authorization", "token "+qa.Get("access_token"))

	var r *http.Response
	r, err = client.Do(req)
	Check(err)

	defer r.Body.Close()

	body, err = io.ReadAll(r.Body)
	Check(err)

	var m *struct {
		Login string "json:\"login\""
	}

	m, err = ParseUserJSON(string(body))
	Check(err)

	var homedir string
	homedir, err = os.UserHomeDir()
	Check(err)

	var exists bool

	if exists, err = PathExists(homedir + "/cakecutter"); err != nil {
		Check(err)
	}

	if !exists {
		err = os.Mkdir(homedir+"/cakecutter", 0755)
		Check(err)
		b := []byte(qa.Get("access_token"))
		err = os.WriteFile(homedir+"/cakecutter/oauth", b, 0644)
		Check(err)
		fmt.Println(Colorize("green", "Logged in as "+m.Login))
		return
	}

	b := []byte(qa.Get("access_token"))
	err = os.WriteFile(homedir+"/cakecutter/oauth", b, 0644)
	Check(err)
	fmt.Println(Colorize("green", "Logged in as "+m.Login))

}

func LoggedIn() (*string, bool, error) {

	loggedIn := false
	var user string
	var yes bool

	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, loggedIn, err
	}

	yes, err = PathExists(homedir + "/cakecutter/oauth")
	if err != nil {
		return nil, loggedIn, err
	}

	if yes {
		var (
			content []byte
			r       *http.Response
			body    []byte
			m       *struct {
				Login string "json:\"login\""
			}
		)

		content, err = os.ReadFile(homedir + "/cakecutter/oauth")
		if err != nil {
			return nil, loggedIn, err
		}

		client := &http.Client{}
		req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
		req.Header.Set("Authorization", "token "+string(content))

		r, err = client.Do(req)
		if err != nil {
			return nil, loggedIn, err
		}

		defer r.Body.Close()

		body, err = io.ReadAll(r.Body)
		if err != nil {
			return nil, loggedIn, err
		}

		m, err = ParseUserJSON(string(body))
		if err != nil {
			return nil, loggedIn, err
		}

		loggedIn = true
		user = m.Login

	}

	return &user, loggedIn, nil

}

func Logout() error {

	var sure bool

	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Are you sure you want to logout? [y/n] ")

	if scanner.Scan() {
		if scanner.Text() != "" {
			__r := scanner.Text()

			if strings.HasPrefix(strings.ToLower(__r), "y") {
				sure = true
			} else {
				sure = false
			}
		}
	}

	if !sure {
		fmt.Println(Colorize("blue", "\nAborted"))
		return nil
	} else {
		err = os.Remove(homedir + "/cakecutter/oauth")
		if err != nil {
			return err
		}
		fmt.Println(Colorize("green", "Logged out successfully"))
	}

	return nil

}
