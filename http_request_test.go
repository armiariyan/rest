package rest

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// implements io.Reader
type nopReader struct {
	err error
}

func (r *nopReader) Read(p []byte) (n int, err error) {
	return len(p), r.err
}

// implements io.ReadCloser
type nopCloser struct {
	io.Reader
	err error
}

func (n *nopCloser) Close() error {
	return n.err
}

// noopCloser returns a ReadCloser with a no-op Close method wrapping
// the provided Reader r.
func noopCloser(r io.Reader, err error) io.ReadCloser {
	return &nopCloser{
		Reader: r,
		err:    err,
	}
}

type mockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var doFuncMock = func(body []byte, err error) func(req *http.Request) (*http.Response, error) {
	return func(req *http.Request) (*http.Response, error) {
		// do whatever you want
		body := noopCloser(bytes.NewReader(body), nil)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       body,
		}, err
	}
}

func (m *mockClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	// just in case you want default correct return value
	return &http.Response{}, nil
}

func TestDefaultClient(t *testing.T) {
	convey.Convey("Test default client", t, func() {
		convey.Convey("When *http.Client is nil, should return panic", func() {
			defer func() {
				if r := recover(); r == nil {
					// function must be panic
					t.Errorf("It must be panic since *http.Client is nil, but it's not!")
					return
				}
			}()

			_, _ = DefaultClient(nil)
		})

		convey.Convey("When *http.Client is not nil, it must return the instance", func() {
			testClient := &mockClient{
				DoFunc: doFuncMock(nil, nil),
			}

			client, err := DefaultClient(testClient)
			convey.So(client, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})

		convey.Convey("When using one options, and error", func() {
			testClient := &mockClient{
				DoFunc: doFuncMock(nil, nil),
			}

			var opt Option = func(requester *DefaultHttpRequester) error {
				return fmt.Errorf("error")
			}

			client, err := DefaultClient(testClient, opt)
			convey.So(client, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
	})
}

func TestDefaultHttpRequesterHook(t *testing.T) {
	convey.Convey("Test AddHook", t, func() {
		convey.Convey("Skip hook if nil, but call it func if not nil", func() {
			// this must call before and after hook, and call the function if it exist in slice of hook
			// skip it if in slice contains null (nil) value
			req := &DefaultHttpRequester{
				hook: []Hook{nil, new(NoopHook)},
			}

			req.beforeHook(context.Background(), HookData{})
			req.afterHook(context.Background(), HookData{})
		})
	})
}

func TestDefaultHttpRequesterGet(t *testing.T) {
	convey.Convey("Test Get", t, func() {

		convey.Convey("Get should be success", func() {
			testClient := &mockClient{
				DoFunc: doFuncMock([]byte(`hello world`), nil),
			}

			client, err := DefaultClient(testClient)
			convey.So(client, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)

			resp, err := client.Get(context.Background(), "", "http://example.com/", http.Header{})
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestDefaultHttpRequesterPost(t *testing.T) {
	convey.Convey("Test Post", t, func() {

		convey.Convey("Post should be success", func() {
			testClient := &mockClient{
				DoFunc: doFuncMock([]byte(`hello world`), nil),
			}

			client, err := DefaultClient(testClient)
			convey.So(client, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)

			resp, err := client.Post(context.Background(), "", "http://example.com/", http.Header{}, nil)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestDefaultHttpRequesterPut(t *testing.T) {
	convey.Convey("Test Put", t, func() {

		convey.Convey("Put should be success", func() {
			testClient := &mockClient{
				DoFunc: doFuncMock([]byte(`hello world`), nil),
			}

			client, err := DefaultClient(testClient)
			convey.So(client, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)

			resp, err := client.Put(context.Background(), "", "http://example.com/", http.Header{}, nil)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})

	})
}

func TestDefaultHttpRequesterPatch(t *testing.T) {
	convey.Convey("Test Patch", t, func() {

		convey.Convey("Patch should be success", func() {
			testClient := &mockClient{
				DoFunc: doFuncMock([]byte(`hello world`), nil),
			}

			client, err := DefaultClient(testClient)
			convey.So(client, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)

			resp, err := client.Patch(context.Background(), "", "http://example.com/", http.Header{}, nil)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})

	})
}

func TestDefaultHttpRequesterDelete(t *testing.T) {
	convey.Convey("Test Delete", t, func() {

		convey.Convey("Delete should be success", func() {
			testClient := &mockClient{
				DoFunc: doFuncMock([]byte(`hello world`), nil),
			}

			client, err := DefaultClient(testClient)
			convey.So(client, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)

			resp, err := client.Delete(context.Background(), "", "http://example.com/", http.Header{}, nil)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})

	})
}

func TestDefaultHttpRequesterCall(t *testing.T) {
	convey.Convey("Test call", t, func() {
		testClient := &mockClient{
			DoFunc: doFuncMock([]byte(`hello world`), nil),
		}

		client := DefaultHttpRequester{
			client: testClient,
			hook:   nil,
		}
		convey.So(client, convey.ShouldNotBeNil)

		convey.Convey("Error url.Parse since it contain new line", func() {
			// to see what causing url.Parse return error, see https://golang.org/doc/go1.12#net/url
			resp, err := client.call(context.Background(), http.MethodPost, "", "http://example.com\n", http.Header{}, nil)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})

		convey.Convey("Post should be success", func() {
			testClient.DoFunc = doFuncMock([]byte(`{}`), nil)

			resp, err := client.call(context.Background(), http.MethodPost, "", "http://example.com/", http.Header{}, nil)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})

		convey.Convey("Error client.Do request", func() {
			testClient.DoFunc = doFuncMock([]byte(`hello world`), fmt.Errorf("error"))

			resp, err := client.call(context.Background(), http.MethodPost, "", "http://example.com/", http.Header{}, nil)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})

		convey.Convey("Error response is nil", func() {
			testClient.DoFunc = func(req *http.Request) (*http.Response, error) {
				return nil, nil
			}

			resp, err := client.call(context.Background(), http.MethodPost, "", "http://example.com/", http.Header{}, nil)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})

		convey.Convey("Error response is nil and error returned", func() {
			testClient.DoFunc = func(req *http.Request) (*http.Response, error) {
				return nil, fmt.Errorf("error return when response nil")
			}

			resp, err := client.call(context.Background(), http.MethodPost, "", "http://example.com/", http.Header{}, nil)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})

		convey.Convey("Error response is nil and HTTP timeout returned", func() {
			testClient.DoFunc = func(req *http.Request) (*http.Response, error) {
				var err = errors.New("error from nowhere")
				return nil, fmt.Errorf(err.Error() + " (Client.Timeout exceeded while awaiting headers)")
			}

			resp, err := client.call(context.Background(), http.MethodPost, "", "http://example.com/", http.Header{}, nil)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldResemble, ErrHttpTimeout)
		})

		convey.Convey("Error body close, but not return error in function (only in defer mode)", func() {
			// test this line:
			// if err := resp.Body.Close(); err != nil {

			testClient.DoFunc = func(req *http.Request) (*http.Response, error) {
				// do whatever you want
				body := noopCloser(bytes.NewReader([]byte("hello world")), fmt.Errorf("error closing body"))
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       body,
				}, nil
			}

			resp, err := client.call(context.Background(), http.MethodPost, "", "http://example.com/", http.Header{}, nil)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldBeNil)
		})

		convey.Convey("Error on ioutil.ReadAll body, but success closing body", func() {
			testClient.DoFunc = func(req *http.Request) (*http.Response, error) {
				body := noopCloser(
					&nopReader{
						err: fmt.Errorf("error read body"),
					},
					fmt.Errorf("error closing body"),
				)

				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       body,
				}, nil
			}

			resp, err := client.call(context.Background(), http.MethodPost, "", "http://example.com/", http.Header{}, nil)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})

	})
}
