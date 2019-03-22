package pkg

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSaveToCSV(t *testing.T) {
	Convey("Test SaveToCSV", t, func() {
		analyzer := NewAnalyzer()
		err := analyzer.SaveToCSV()
		So(err, ShouldBeNil)
	})
}
