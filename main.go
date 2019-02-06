package kotrcli

import (
	"errors"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	ini "gopkg.in/ini.v1"
	cli "gopkg.in/urfave/cli.v1"
)

func getTokens() (string, string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", "", err
	}
	ini, err := ini.Load(filepath.Join(home, ".kot", "config"))
	if err != nil {
		return "", "", err
	}
	sec, err := ini.GetSection("default")
	if err != nil {
		return "", "", err
	}
	if !sec.HasKey("user_token") {
		return "", "", errors.New("please define `user_token`")
	}
	if !sec.HasKey("token") {
		return "", "", errors.New("please define `user_token`")
	}
	userToken, _ := sec.GetKey("user_token")
	token, _ := sec.GetKey("token")
	return userToken.String(), token.String(), nil

}

func initApp() (*cli.App, error) {
	userToken, token, err := getTokens()
	if err != nil {
		return nil, err
	}
	app := cli.NewApp()
	app.Name = "kotrcli"
	app.Description = "get rid of your stress at work"
	app.Commands = []cli.Command{
		{
			Name:        "syukkin",
			Aliases:     []string{"s"},
			Usage:       "syukkin",
			Description: "syukkin",
			Action: func(c *cli.Context) error {
				return Dakoku(SYUKKIN, userToken, token)
			},
		},
		{
			Name:        "taikin",
			Aliases:     []string{"t"},
			Usage:       "taikin",
			Description: "taikin",
			Action: func(c *cli.Context) error {
				return Dakoku(TAIKIN, userToken, token)
			},
		},
	}
	return app, nil
}

func mainInternal() int {
	app, err := initApp()
	if err != nil {
		return 1
	}
	err = app.Run(os.Args)
	if err != nil {
		return 1
	}
	return 0
}

func main() {
	os.Exit(mainInternal())
}
