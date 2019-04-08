package pkg

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

// Analyzer represents the data analyzer.
type Analyzer struct {
	OutputFile string
}

// Data represents the response data.
type Data struct {
	RateDetail RateDetail `json:"rateDetail"`
}

// RateDetail contains the reate detail info.
type RateDetail struct {
	Paginator Paginator  `json:"paginator"`
	RateList  []RateData `json:"rateList"`
}

// RateData contains user's rate content.
type RateData struct {
	RateDate    string
	RateContent string
	AuctionSku  string
	Useful      bool
}

// Paginator represents the paginator.
type Paginator struct {
	LastPage int
	Page     int
	Items    int
}

// ItemInfo represents the item search data info.
type ItemInfo struct {
	ShopID       string `json:"shop_id"`
	UserID       string `json:"user_id"`
	ShopTitle    string `json:"shop_title"`
	TotalPage    string `json:"total_page"`
	ShopURL      string `json:"shop_Url"`
	Items        []Item
	CurrentPage  string `json:"current_page"`
	PageSize     string `json:"page_size"`
	TotalResults string `json:"total_results"`
}

// Item represents the item info.
type Item struct {
	ItemID            int64  `json:"item_id"`
	Title             string `json:"title"`
	Img               string `json:"img"`
	Sold              string `json:"sold"`
	Quantity          int
	TotalSoldQuantity int `json:"totalSoldQuantity"`
	URL               string
	Price             string
}

// NewAnalyzer returns analyzer
func NewAnalyzer(output string) *Analyzer {
	return &Analyzer{
		OutputFile: output,
	}
}

// ParseRate parses rate data from JSON data.
func ParseRate(data []byte) (*RateDetail, error) {
	jsonData := &Data{}
	err := json.Unmarshal(data, jsonData)
	if err != nil {
		return nil, err
	}
	return &jsonData.RateDetail, nil
}

// ParseItem parses item data from JSON data.
func ParseItem(data []byte) (*ItemInfo, error) {
	item := &ItemInfo{}
	err := json.Unmarshal(data, item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// SaveItemToCSV saves item data to CSV file.
func (a *Analyzer) SaveItemToCSV(itemData *ItemInfo) error {
	records := [][]string{}
	for _, item := range itemData.Items {
		records = append(records, []string{
			strconv.FormatInt(item.ItemID, 10),
			itemData.UserID,
			item.Title, item.Img, item.Price, item.Sold,
			strconv.Itoa(item.Quantity),
			strconv.Itoa(item.TotalSoldQuantity),
			item.URL,
		})
	}
	return saveToCSV(a.OutputFile, records)
}

// SaveRateToCSV saves rate data to CSV file.
func (a *Analyzer) SaveRateToCSV(rate *RateDetail) error {
	records := [][]string{}
	for _, record := range rate.RateList {
		records = append(records, []string{
			record.RateDate, record.RateContent, record.AuctionSku,
		})
	}
	return saveToCSV(a.OutputFile, records)
}

func saveToCSV(filename string, records [][]string) error {
	csvFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644)
	if err != nil {
		log.Errorf("create output file(%s) failed: %v", filename, err)
		return err
	}
	defer csvFile.Close()
	// set UTF-8 BOM
	csvFile.WriteString("\xEF\xBB\xBF")
	csvWriter := csv.NewWriter(csvFile)
	return csvWriter.WriteAll(records)
}
