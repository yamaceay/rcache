package main

import (
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	ptr *http.Request
}

func (req *Request) Res() string {
	resp, err := http.DefaultClient.Do(req.ptr)
	if err != nil {
		fmt.Printf("server returns %d error: %s", resp.StatusCode, err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	return string(bodyBytes)
}

func Req(address string, method string, key string, value string) *Request {
	fullPath := fmt.Sprintf("%s/%s", address, method)

	req, _ := http.NewRequest("GET", fullPath, nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	q := req.URL.Query()
	if key != "" {
		q.Add("key", key)
	}
	if value != "" {
		q.Add("value", value)
	}
	req.URL.RawQuery = q.Encode()
	reqInternal := Request{req}
	return &reqInternal
}
