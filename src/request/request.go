package request

import (
	"errors"
	"io"
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

func Request(url string, option RequestOption) ([]byte, error) {
	var defaultResult []byte
	method := option.Method

	if len(url) == 0 {
		return defaultResult, ErrRequestEmptyUrl
	}

	parsedUrl, err := u.Parse(url)
	if err != nil {
		return defaultResult, ErrRequestInvalidUrl
	}
	if parsedUrl.Scheme == "" {
		parsedUrl.Scheme = "http"
	}

	req, err := http.NewRequest(method, parsedUrl.String(), nil)
	if err != nil {
		println(err.Error())
		return defaultResult, ErrRequestFailedToCreate
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		println(err.Error())
		return defaultResult, ErrRequestProtocol
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		println(err.Error())
		return defaultResult, ErrRequestParseResBody
	}

	return body, nil
}
