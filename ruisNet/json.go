package ruisNet

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

func NewRequestJson(method, urls string, body interface{}) (*http.Request, error) {
	bts, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, urls, bytes.NewBuffer(bts))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	return req, nil
}
func DoHttpJson(method, urls string, body interface{}, timeout ...time.Duration) (*http.Response, error) {
	req, err := NewRequestJson(method, urls, body)
	if err != nil {
		return nil, err
	}
	cli := &http.Client{}
	if len(timeout) > 0 {
		cli.Timeout = timeout[0]
	}
	return cli.Do(req)
}
