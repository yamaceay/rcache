package rcache

import (
	"fmt"
	"io"
	"net/http"
)

func Send(address string, method string, key string, value string) (string, error) {
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

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("server returns %d error: %s", resp.StatusCode, err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	return string(bodyBytes), nil
}
