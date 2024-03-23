package rest

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestWithCircuitBreaker(t *testing.T) {
	convey.Convey("Test WithCircuitBreaker", t, func() {
		convey.Convey("Should return no error", func() {
			opt := []Option{
				WithCircuitBreaker(CBConfig{}),
			}

			client, err := DefaultClient(new(mockClient), opt...)
			convey.So(client, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestAddHook(t *testing.T) {
	convey.Convey("Test AddHook", t, func() {
		convey.Convey("Should return error when hook is nil", func() {
			opt := []Option{
				AddHook(nil),
			}

			client, err := DefaultClient(new(mockClient), opt...)
			convey.So(client, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})

		convey.Convey("Should return success when hook is not nil", func() {
			opt := []Option{
				AddHook(new(NoopHook)),
			}

			client, err := DefaultClient(new(mockClient), opt...)
			convey.So(client, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
