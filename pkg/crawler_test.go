package pkg

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCrawling(t *testing.T) {
	Convey("Test Crawling", t, func() {
		crawler := &Crawler{}
		err := crawler.Crawling("545414402527", "2635590370")
		So(err, ShouldBeNil)
	})
}
