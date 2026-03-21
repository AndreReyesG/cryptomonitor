package exchanges

import (
	"io"
	"net/http"
	"strings"
	"time"
)

type StubTime struct {
	T time.Time
}

func (s StubTime) Now() time.Time {
	return s.T
}

type StubClient struct {
	resp *http.Response
	err  error
}

func (s *StubClient) Get(url string) (*http.Response, error) {
	return s.resp, s.err
}

func NewStubClient(status int, body string) *StubClient {
	return &StubClient{
		resp: &http.Response{
			StatusCode: status,
			Body:       io.NopCloser(strings.NewReader(body)),
		},
		err: nil,
	}
}
