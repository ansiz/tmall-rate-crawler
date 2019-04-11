// Package pkg contains the crawler implement.
package pkg

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
)

const (
	baseURL = `https://rate.tmall.com/list_detail_rate.htm`
)

// Crawler represents the data crawler manager.
type Crawler struct {
	analyzer *Analyzer
	config   *Config
}

// Config contains the crawler configurations.
type Config struct {
	Cookie             string
	UserAgent          string
	ItemID             string
	SellerID           string
	Output             string
	SortKey            string
	RequestIntervalMin int
	RequestIntervalMax int
}

// NewCrawler creates new crawler.
func NewCrawler(c *Config) *Crawler {
	return &Crawler{
		analyzer: NewAnalyzer(c.Output),
		config: &Config{
			UserAgent:          c.UserAgent,
			Cookie:             c.Cookie,
			ItemID:             c.ItemID,
			SellerID:           c.SellerID,
			RequestIntervalMin: c.RequestIntervalMin,
			RequestIntervalMax: c.RequestIntervalMax,
			SortKey:            c.SortKey,
		},
	}
}

// CrawlItems crawl shop items data.
func (c *Crawler) CrawlItems(shopURL string, page int) error {
	req, err := http.NewRequest("GET", shopURL, nil)
	if err != nil {
		return err
	}
	params := url.Values{}
	params.Add("p", strconv.Itoa(page))
	params.Add("sort", c.config.SortKey)
	req.URL.RawQuery = params.Encode()
	log.Debug("request URI:", req.URL.String())

	req.Header.Add("user-agent", c.config.UserAgent)
	req.Header.Add("cookie", c.config.Cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("request error: %v", err)
		return err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	itemData, err := ParseItem(body)
	if err != nil {
		log.Error("parse data error: %v", err)
		return err
	}
	err = c.analyzer.SaveItemToCSV(itemData)
	if err != nil {
		return err
	}
	cur, _ := strconv.Atoi(itemData.CurrentPage)
	total, _ := strconv.Atoi(itemData.TotalPage)
	log.Infof("Task(shop:%s, has %s items) progress: %d/%d",
		itemData.ShopTitle, itemData.TotalResults, cur, total)
	if cur != total {
		c.randomSleep()
		return c.CrawlItems(shopURL, cur+1)
	}
	return nil
}

// CrawlItemRate crawl the item data.
func (c *Crawler) CrawlItemRate(page int) error {
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return err
	}
	params := url.Values{}
	params.Add("itemId", c.config.ItemID)
	params.Add("sellerId", c.config.SellerID)
	params.Add("currentPage", strconv.Itoa(page))
	req.URL.RawQuery = params.Encode()
	log.Debug("request URI:", req.URL.String())

	req.Header.Add("authority", "rate.tmall.com")
	req.Header.Add("user-agent", c.config.UserAgent)
	req.Header.Add("cookie", c.config.Cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("request error: %v", err)
		return err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// trim response data(json128())
	rateDetail, err := ParseRate(body[11 : len(body)-1])
	if err != nil {
		log.Error("parse data error: %v", err)
		return err
	}
	err = c.analyzer.SaveRateToCSV(rateDetail)
	if err != nil {
		return err
	}
	log.Debugf("Task(item:%s, seller:%s) progress: %d/%d",
		c.config.ItemID, c.config.SellerID, page, rateDetail.Paginator.LastPage)
	if rateDetail.Paginator.Page != rateDetail.Paginator.LastPage {
		c.randomSleep()
		return c.CrawlItemRate(page + 1)
	}
	return nil
}

func (c *Crawler) randomSleep() {
	interval := GenerateRangeNum(c.config.RequestIntervalMin,
		c.config.RequestIntervalMax)
	time.Sleep(time.Duration(interval) * time.Second)
}

// GenerateRangeNum generates random number in specified range.
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}
