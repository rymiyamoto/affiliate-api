package main

import (
	"os"

	"github.com/rymiyamoto/affiliate-api/conf"
	"github.com/rymiyamoto/affiliate-api/dispatch"
	"github.com/rymiyamoto/affiliate-api/util/log"
	"github.com/urfave/cli/v2"
)

func main() {
	if err := conf.InitLog("SCRIPTS"); err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.Name = "affiliate-api-tester"
	app.Usage = "アフィリエイトAPIの動作確認"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "shop",
			Usage:   "対象",
			Value:   "",
			Aliases: []string{"s"},
		},
		&cli.StringFlag{
			Name:    "code",
			Usage:   "商品コード",
			Value:   "",
			Aliases: []string{"c"},
		},
	}
	app.Action = func(c *cli.Context) error {
		s := dispatch.NewAffiliate()

		log.Info("start")

		err := s.Run(c)

		log.Info("finish")

		return err
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
