package main

import (
	"errors"
	"os"

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
		cli.IntFlag{
			Name:  "interval-min",
			Value: 5,
			Usage: "The minion request interval(Second)",
		},
		cli.IntFlag{
			Name:  "interval-max",
			Value: 15,
			Usage: "The max request interval(Second)",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "rate",
			Usage: "fetch specified item's rate data",
			Flags: []cli.Flag{
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
					UserAgent:          c.GlobalString("user-agent"),
					Cookie:             c.GlobalString("cookie"),
					ItemID:             c.String("item"),
					SellerID:           c.String("seller"),
					RequestIntervalMin: c.GlobalInt("interval-min"),
					RequestIntervalMax: c.GlobalInt("interval-max"),
					Output:             c.String("output"),
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
				err := crawler.CrawlItemRate(c.Int("start"))
				if err != nil {
					log.Error("fetch data error: %v", err)
					return err
				}
				return nil
			},
		}, {
			Name:  "item",
			Usage: "fetch shop items data",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "url",
					Value: "",
					Usage: "The shop URL",
				},
				cli.IntFlag{
					Name:  "start",
					Value: 1,
					Usage: "The start page number",
				},
				cli.StringFlag{
					Name:  "sort",
					Value: "d",
					Usage: "The sort rule (d:销量 s:综合 p:价格 oldstarts:新品优先)",
				},
				cli.StringFlag{
					Name:  "output,o",
					Value: "shop-items.csv",
					Usage: "The output file name",
				},
			},
			Action: func(c *cli.Context) error {
				cfg := &pkg.Config{
					UserAgent:          c.GlobalString("user-agent"),
					Cookie:             c.GlobalString("cookie"),
					RequestIntervalMin: c.GlobalInt("interval-min"),
					RequestIntervalMax: c.GlobalInt("interval-max"),
					Output:             c.String("output"),
					SortKey:            c.String("sort"),
				}
				if cfg.Cookie == "" {
					return errors.New("missing required cookie data")
				}
				shopURL := c.String("url")
				if shopURL == "" {
					return errors.New("missing required shop URL")
				}
				crawler := pkg.NewCrawler(cfg)
				err := crawler.CrawlItems(shopURL, c.Int("start"))
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
