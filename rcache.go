package rcache

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Is equal to Send("set", address, "", "")
func Keys(address string) ([]string, error) {
	var keys []string
	if response, err := Send(address, "get", "", ""); err != nil {
		return []string{}, fmt.Errorf("unable to fetch all keys: %s", err)
	} else if err := json.Unmarshal([]byte(response), &keys); err != nil {
		return []string{}, fmt.Errorf("cannot unmarshal keys: %s", err)
	}
	return keys, nil
}

// Is equal to Send("get", address, key, "")
func Get(address string, key string) (string, error) {
	return Send(address, "get", key, "")
}

// Is equal to Send("set", address, key, value)
func Set(address string, key string, value string) (string, error) {
	return Send("set", address, key, value)
}

// Creates a HTTP request and returns the response string
//
// Arguments:
//  * address string
//  * method string
//  * key string
//  * value string
// Returns:
// 	* response string
// 	* err error
func Send(method string, address string, key string, value string) (string, error) {
	fullPath := fmt.Sprintf("%s/%s", address, method)

	req, err := http.NewRequest("GET", fullPath, nil)
	if err != nil {
		return "", fmt.Errorf("http request failed: %s", err)
	}
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

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("body cannot be read: %s", err)
	}
	return string(bodyBytes), nil
}
