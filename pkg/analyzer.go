package pkg

import (
	"encoding/csv"
	"encoding/json"
	"os"
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

// NewAnalyzer returns analyzer
func NewAnalyzer(output string) *Analyzer {
	return &Analyzer{
		OutputFile: output,
	}
}

// SaveToCSV saves data to CSV file.
func (a *Analyzer) SaveToCSV(rate *RateDetail) error {
	records := [][]string{}
	for _, record := range rate.RateList {
		records = append(records, []string{
			record.RateDate, record.RateContent, record.AuctionSku,
		})
	}
	csvFile, err := os.OpenFile("result.csv",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer csvFile.Close()
	// set UTF-8 BOM
	csvFile.WriteString("\xEF\xBB\xBF")
	csvWriter := csv.NewWriter(csvFile)
	return csvWriter.WriteAll(records)
}

// ParseRate parse data from JSON data.
func ParseRate(data []byte) (*RateDetail, error) {
	jsonData := &Data{}
	err := json.Unmarshal(data, jsonData)
	if err != nil {
		return nil, err
	}
	return &jsonData.RateDetail, nil
}
