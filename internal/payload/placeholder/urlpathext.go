package placeholder

import (
	"net/http"
	"net/url"
)

type URLFileExt struct {
	name string
}

var DefaultURLFileExt = URLFileExt{name: "URLFileExt"}

var _ Placeholder = (*URLFileExt)(nil)

func (p URLFileExt) GetName() string {
	return p.name
}

func (p URLFileExt) CreateRequest(requestURL, payload string) (*http.Request, error) {
	reqURL, err := url.Parse(requestURL)
	if err != nil {
		return nil, err
	}

	urlWithPayload := reqURL.String()
	for i := len(urlWithPayload) - 1; i >= 0; i-- {
		if urlWithPayload[i] != '/' {
			urlWithPayload = urlWithPayload[:i+1]
			break
		}
	}
	urlWithPayload += "/index." + payload

	req, err := http.NewRequest("GET", urlWithPayload, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
