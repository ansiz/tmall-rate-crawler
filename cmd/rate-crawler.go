package main

import (
	"errors"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/ansiz/tmall-rate-crawler/pkg"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Tmall rate crawler"
	app.Usage = "Crawl rate data from Tmall"
	app.Version = "0.1.0"
	app.Before = appInit
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "run in verbose mode",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "fetch",
			Usage: "fetch specified item's rate data",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "user-agent",
					Value: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36",
					Usage: "The HTTP request user-agent header",
				},
				cli.StringFlag{
					Name:  "cookie,c",
					Value: "",
					Usage: "The cookie data",
				},
				cli.StringFlag{
					Name:  "item,i",
					Value: "",
					Usage: "The item ID",
				},
				cli.StringFlag{
					Name:  "seller,s",
					Value: "",
					Usage: "The seller ID",
				},
				cli.IntFlag{
					Name:  "interval",
					Value: 3,
					Usage: "The request interval(Second)",
				},
				cli.IntFlag{
					Name:  "start",
					Value: 1,
					Usage: "The start page number",
				},
				cli.StringFlag{
					Name:  "output,o",
					Value: "result.csv",
					Usage: "The output file name",
				},
			},
			Action: func(c *cli.Context) error {
				cfg := &pkg.Config{
					UserAgent:       c.String("user-agent"),
					Cookie:          c.String("cookie"),
					ItemID:          c.String("item"),
					SellerID:        c.String("seller"),
					RequestInterval: time.Second * time.Duration(c.Int("interval")),
				}
				if cfg.Cookie == "" {
					return errors.New("missing required cookie data")
				}
				if cfg.ItemID == "" {
					return errors.New("missing required item id")
				}
				if cfg.SellerID == "" {
					return errors.New("missing required seller id")
				}
				crawler := pkg.NewCrawler(cfg)
				err := crawler.Crawling(c.Int("start"))
				if err != nil {
					log.Error("fetch data error: %v", err)
					return err
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func appInit(ctx *cli.Context) error {
	if ctx.Bool("verbose") {
		log.SetLevel(log.DebugLevel)
	}
	return nil
}
