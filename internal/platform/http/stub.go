package http

import (
	"io"
	"net/http"
	"strings"
)

type Stub struct {
	resp *http.Response
	err  error
}

func (s *Stub) Get(url string) (*http.Response, error) {
	return s.resp, s.err
}

func NewStub(status int, body string) *Stub {
	return &Stub{
		resp: &http.Response{
			StatusCode: status,
			Body:       io.NopCloser(strings.NewReader(body)),
		},
		err: nil,
	}
}
