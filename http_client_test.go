package rest

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestNewHttpClient(t *testing.T) {
	convey.Convey("Test New HTTP Client", t, func() {
		convey.Convey("When Transport.TLSHandshakeTimeout or Client.Timeout <= 0", func() {
			conf := Config{}
			netClient := NewHttpClient(conf)
			convey.So(netClient, convey.ShouldNotBeNil)
		})
	})
}
