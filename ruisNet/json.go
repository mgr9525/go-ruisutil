package ruisNet

import (
	"bytes"
	"encoding/json"
	"net/http"
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
