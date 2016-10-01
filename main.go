package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"net/http"
	"time"
	"net/url"
	"errors"
	"strings"
)

var RedirectAttemptedError = errors.New("redirect")

func main() {
	app := cli.NewApp()
	app.Name = "lgtm"
	app.Version = "1.0.0"
	app.Usage = "get http://lgtm.in/g image url"
	app.Action = func(c *cli.Context) error {

		location, err := getLocation("http://lgtm.in/g")
		if err != nil {
			return err
		}

		new_url := strings.Replace(location[0], "/i/", "/p/", 1)

		lgtm_url, err := getLocation(new_url)
		if err != nil {
			return err
		}

		fmt.Println(lgtm_url[0])
		return nil
	}

	app.Run(os.Args)
}

func getLocation(target_url string) ([]string, error) {
	client := &http.Client{
		Timeout: time.Duration(3) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return RedirectAttemptedError
		},
	}
	resp, err := client.Head(target_url)
	defer resp.Body.Close()
	if urlError, ok := err.(*url.Error); ok && urlError.Err == RedirectAttemptedError {
		return resp.Header["Location"], nil
	}

	return nil, err
}
