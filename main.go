package main

import (
	"fmt"
	"os"

	"github.com/rymiyamoto/affiliate-api/conf"
	"github.com/rymiyamoto/affiliate-api/log"
	"github.com/urfave/cli"
)

func init() {
	conf.InitLog("SCRIPTS")
}

func main() {
	app := cli.NewApp()
	app.Name = "push"
	app.Usage = "プッシュ通知メッセージを送信します"
	app.Commands = commands
	err := app.Run(os.Args)
	if err != nil {
		log.Errorf("処理が失敗しました: %v", err)
		os.Exit(1)
	}
}

var commands = []cli.Command{
	{
		Name:  "rakuten",
		Usage: "楽天商品検索API",
		Flags: []cli.Flag{
			cli.StringFlag{
				Required: false,
				Name:     "date,d",
				Usage:    "基準日 `yyyy-mm-dd` (default today",
			},
		},
		Action: Hoge,
	},
	{
		Name:  "yahoo",
		Usage: "Yahoo商品検索API",
		Flags: []cli.Flag{
			cli.StringFlag{
				Required: false,
				Name:     "date,d",
				Usage:    "基準日 `yyyy-mm-dd` (default today",
			},
		},
		Action: Hoge,
	},
}

func Hoge() {
	fmt.Println("hello world")
}
