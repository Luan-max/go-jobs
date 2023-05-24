package helpers

import (
	"bytes"
	"errors"
	"net/http"
)

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)

func JSONRequest(method HttpMethod, url string, body []byte, headers map[string]string) (*http.Request, error) {
	switch method {
	case GET, POST, PUT, DELETE:
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, errors.New(err.Error())
		}

		for key, value := range headers {
			req.Header.Set(key, value)
		}

		return req, nil
	default:
		return nil, errors.New(string(method) + "is not a valid HTTP method")
	}
}
