package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSaveToCSV(t *testing.T) {
	Convey("Test SaveToCSV", t, func() {
		testfile := "/tmp/analyzer.test"
		analyzer := NewAnalyzer("/tmp/analyzer.test")
		err := analyzer.SaveToCSV(&RateDetail{})
		So(err, ShouldBeNil)
		os.Remove(testfile)
	})
}

func TestParser(t *testing.T) {
	Convey("Test Item Parser", t, func() {
		data, err := ioutil.ReadFile("../test/item-data.json")
		So(err, ShouldBeNil)
		So(data, ShouldNotBeNil)
		itemInfo, err := ParseItem(data)
		So(err, ShouldBeNil)
		So(itemInfo, ShouldNotBeNil)
		fmt.Printf("---%#v", itemInfo)
	})
}
