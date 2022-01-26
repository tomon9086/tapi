package request

import (
	"errors"
	"net/http"
	u "net/url"
)

type RequestOption struct {
	Method string
}

var (
	ErrRequestEmptyUrl       = errors.New("empty url")
	ErrRequestInvalidUrl     = errors.New("invalid url")
	ErrRequestFailedToCreate = errors.New("failed to create request")
	ErrRequestProtocol       = errors.New("protocol error")
	ErrRequestParseResBody   = errors.New("failed to parse response body")
)

func Request(url string, option RequestOption) (*http.Request, *http.Response, error) {
	method := option.Method

	if len(url) == 0 {
		return nil, nil, ErrRequestEmptyUrl
	}

	parsedUrl, err := u.Parse(url)
	if err != nil {
		return nil, nil, ErrRequestInvalidUrl
	}
	if parsedUrl.Scheme == "" {
		parsedUrl.Scheme = "http"
	}

	req, err := http.NewRequest(method, parsedUrl.String(), nil)
	if err != nil {
		println(err.Error())
		return nil, nil, ErrRequestFailedToCreate
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		println(err.Error())
		return nil, nil, ErrRequestProtocol
	}

	return req, res, nil
}
