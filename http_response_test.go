package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestHttpResponseTo(t *testing.T) {
	convey.Convey("HttpResponse.To", t, func() {
		convey.Convey("HttpResponse is nil", func() {
			resp := &HttpResponse{}
			resp = nil

			var data interface{}
			err := resp.To(context.Background(), json.Unmarshal, &data)
			convey.So(err, convey.ShouldResemble, fmt.Errorf("h is nil"))
		})

		convey.Convey("Response body is nil", func() {
			resp := &HttpResponse{
				RespBody: nil,
			}

			var data interface{}
			err := resp.To(context.Background(), json.Unmarshal, &data)
			convey.So(err, convey.ShouldResemble, fmt.Errorf("response body is nil"))
		})

		convey.Convey("Success marshalling to struct", func() {
			resp := &HttpResponse{
				RespBody: []byte(`{"foo": "bar"}`),
			}

			var want = map[string]string{
				"foo": "bar",
			}

			var data map[string]string
			err := resp.To(context.Background(), json.Unmarshal, &data)

			convey.So(data, convey.ShouldResemble, want)
			convey.So(err, convey.ShouldBeNil)
		})

	})
}

func TestHttpResponseToJson(t *testing.T) {
	convey.Convey("HttpResponse.ToJson", t, func() {
		convey.Convey("Success marshalling to struct", func() {
			resp := &HttpResponse{
				RespBody: []byte(`{"amount": 1234}`),
			}

			var want = map[string]int64{
				"amount": 1234,
			}

			var data map[string]int64
			err := resp.ToJson(context.Background(), &data)

			convey.So(data, convey.ShouldResemble, want)
			convey.So(err, convey.ShouldBeNil)
		})

	})
}
