package rest

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

var wantResp = map[string]interface{}{
	"foo": "bar",
}

func TestNewMock(t *testing.T) {
	convey.Convey("New httpclient.Mock", t, func() {
		convey.Convey("Return mock object", func() {
			client := NewMock()
			convey.So(client, convey.ShouldNotBeNil)
		})
	})
}

func TestMockGet(t *testing.T) {
	convey.Convey("New httpclient.Mock", t, func() {
		convey.Convey("When call Get then return not HttpResponse object", func() {
			client := NewMock()
			convey.So(client, convey.ShouldNotBeNil)

			ctx := context.Background()
			client.On("Get", ctx, mock.Anything, "/", http.Header{}).
				Return(wantResp, nil)

			resGet, errGet := client.Get(ctx, "", "/", http.Header{})

			// should return empty HttpResponse on mock
			convey.So(resGet, convey.ShouldResemble, HttpResponse{})
			convey.So(errGet, convey.ShouldNotBeNil)
			convey.So(errGet, convey.ShouldResemble, fmt.Errorf("not HttpResponse type"))
		})

		convey.Convey("When call Get then return expected", func() {
			client := NewMock()
			convey.So(client, convey.ShouldNotBeNil)

			want := HttpResponse{
				RespBody: []byte("hello"),
				CURL:     "CURL -X GET /",
				Raw:      ResponseRaw{},
			}

			ctx := context.Background()
			client.On("Get", ctx, mock.Anything, "/", http.Header{}).
				Return(want, nil)

			resGet, errGet := client.Get(ctx, "", "/", http.Header{})
			convey.So(resGet, convey.ShouldResemble, want)
			convey.So(errGet, convey.ShouldBeNil)
		})
	})
}

func TestMockPost(t *testing.T) {
	convey.Convey("New httpclient.Mock", t, func() {
		convey.Convey("When call Post then return not HttpResponse object", func() {
			client := NewMock()
			convey.So(client, convey.ShouldNotBeNil)

			ctx := context.Background()
			client.On("Post", ctx, mock.Anything, "/", http.Header{}, []byte(nil)).
				Return(wantResp, nil)

			resPost, errPost := client.Post(ctx, "", "/", http.Header{}, []byte(nil))

			// should return empty HttpResponse on mock
			convey.So(resPost, convey.ShouldResemble, HttpResponse{})
			convey.So(errPost, convey.ShouldNotBeNil)
			convey.So(errPost, convey.ShouldResemble, fmt.Errorf("not HttpResponse type"))
		})

		convey.Convey("When call Post then return expected", func() {
			client := NewMock()
			convey.So(client, convey.ShouldNotBeNil)

			want := HttpResponse{
				RespBody: []byte("hello"),
				CURL:     "CURL -X POST /",
				Raw:      ResponseRaw{},
			}

			ctx := context.Background()
			client.On("Post", ctx, mock.Anything, "/", http.Header{}, []byte(nil)).
				Return(want, nil)

			resPost, errPost := client.Post(ctx, "", "/", http.Header{}, []byte(nil))
			convey.So(resPost, convey.ShouldResemble, want)
			convey.So(errPost, convey.ShouldBeNil)
		})
	})
}

func TestMockPut(t *testing.T) {
	convey.Convey("New httpclient.Mock", t, func() {
		convey.Convey("When call Put then return not HttpResponse object", func() {
			client := NewMock()
			convey.So(client, convey.ShouldNotBeNil)

			ctx := context.Background()
			client.On("Put", ctx, mock.Anything, "/", http.Header{}, []byte(nil)).
				Return(wantResp, nil)

			resPut, errPut := client.Put(ctx, "", "/", http.Header{}, []byte(nil))

			// should return empty HttpResponse on mock
			convey.So(resPut, convey.ShouldResemble, HttpResponse{})
			convey.So(errPut, convey.ShouldNotBeNil)
			convey.So(errPut, convey.ShouldResemble, fmt.Errorf("not HttpResponse type"))
		})

		convey.Convey("When call Put then return expected", func() {
			client := NewMock()
			convey.So(client, convey.ShouldNotBeNil)

			want := HttpResponse{
				RespBody: []byte("hello"),
				CURL:     "CURL -X PUT /",
				Raw:      ResponseRaw{},
			}

			ctx := context.Background()
			client.On("Put", ctx, mock.Anything, "/", http.Header{}, []byte(nil)).
				Return(want, nil)

			resPut, errPut := client.Put(ctx, "", "/", http.Header{}, []byte(nil))
			convey.So(resPut, convey.ShouldResemble, want)
			convey.So(errPut, convey.ShouldBeNil)
		})
	})
}

func TestMockPatch(t *testing.T) {
	convey.Convey("New httpclient.Mock", t, func() {
		convey.Convey("When call Patch then return not HttpResponse object", func() {
			client := NewMock()
			convey.So(client, convey.ShouldNotBeNil)

			ctx := context.Background()
			client.On("Patch", ctx, mock.Anything, "/", http.Header{}, []byte(nil)).
				Return(wantResp, nil)

			resPut, errPut := client.Patch(ctx, "", "/", http.Header{}, []byte(nil))

			// should return empty HttpResponse on mock
			convey.So(resPut, convey.ShouldResemble, HttpResponse{})
			convey.So(errPut, convey.ShouldNotBeNil)
			convey.So(errPut, convey.ShouldResemble, fmt.Errorf("not HttpResponse type"))
		})

		convey.Convey("When call Patch then return expected", func() {
			client := NewMock()
			convey.So(client, convey.ShouldNotBeNil)

			want := HttpResponse{
				RespBody: []byte("hello"),
				CURL:     "CURL -X PATCH /",
				Raw:      ResponseRaw{},
			}

			ctx := context.Background()
			client.On("Patch", ctx, mock.Anything, "/", http.Header{}, []byte(nil)).
				Return(want, nil)

			resPut, errPut := client.Patch(ctx, "", "/", http.Header{}, []byte(nil))
			convey.So(resPut, convey.ShouldResemble, want)
			convey.So(errPut, convey.ShouldBeNil)
		})
	})
}

func TestMockDelete(t *testing.T) {
	convey.Convey("New httpclient.Mock", t, func() {
		convey.Convey("When call Delete then return not HttpResponse object", func() {
			client := NewMock()
			convey.So(client, convey.ShouldNotBeNil)

			ctx := context.Background()
			client.On("Delete", ctx, mock.Anything, "/", http.Header{}, []byte(nil)).
				Return(wantResp, nil)

			resPut, errPut := client.Delete(ctx, "", "/", http.Header{}, []byte(nil))

			// should return empty HttpResponse on mock
			convey.So(resPut, convey.ShouldResemble, HttpResponse{})
			convey.So(errPut, convey.ShouldNotBeNil)
			convey.So(errPut, convey.ShouldResemble, fmt.Errorf("not HttpResponse type"))
		})

		convey.Convey("When call Delete then return expected", func() {
			client := NewMock()
			convey.So(client, convey.ShouldNotBeNil)

			want := HttpResponse{
				RespBody: []byte("hello"),
				CURL:     "CURL -X DELETE /",
				Raw:      ResponseRaw{},
			}

			ctx := context.Background()
			client.On("Delete", ctx, mock.Anything, "/", http.Header{}, []byte(nil)).
				Return(want, nil)

			resPut, errPut := client.Delete(ctx, "", "/", http.Header{}, []byte(nil))
			convey.So(resPut, convey.ShouldResemble, want)
			convey.So(errPut, convey.ShouldBeNil)
		})
	})
}
