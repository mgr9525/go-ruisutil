package ruisNet

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
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
func DoHttpJsons(method, urls string, body interface{}, timeout ...time.Duration) (int, []byte, error) {
	req, err := NewRequestJson(method, urls, body)
	if err != nil {
		return 0, nil, err
	}
	cli := &http.Client{}
	if len(timeout) > 0 {
		cli.Timeout = timeout[0]
	}
	res, err := cli.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()
	bts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, nil, err
	}
	return res.StatusCode, bts, err
}
func DoHttpJsonObj(method, urls string, body, rets interface{}, timeout ...time.Duration) error {
	if rets == nil {
		return errors.New("rets is nil")
	}
	req, err := NewRequestJson(method, urls, body)
	if err != nil {
		return err
	}
	cli := &http.Client{}
	if len(timeout) > 0 {
		cli.Timeout = timeout[0]
	}
	res, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	bts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New(string(bts))
	}
	err = json.Unmarshal(bts, rets)
	if err != nil {
		return err
	}
	return nil
}

func NewRequest(method, urls string, body interface{}) (*http.Request, error) {
	var buf = bytes.NewBuffer([]byte{})
	if body != nil {
		switch body.(type) {
		case []byte:
			buf = bytes.NewBuffer(body.([]byte))
		case string:
			buf = bytes.NewBuffer([]byte(body.(string)))
		default:
			bts, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			buf = bytes.NewBuffer(bts)
		}
	}
	return http.NewRequest(method, urls, buf)
}
func DoHttp(method, urls string, body interface{}, timeout ...time.Duration) (*http.Response, error) {
	req, err := NewRequest(method, urls, body)
	if err != nil {
		return nil, err
	}
	cli := &http.Client{}
	if len(timeout) > 0 {
		cli.Timeout = timeout[0]
	}
	return cli.Do(req)
}
func DoHttps(method, urls string, body interface{}, timeout ...time.Duration) (int, []byte, error) {
	req, err := NewRequest(method, urls, body)
	if err != nil {
		return 0, nil, err
	}
	cli := &http.Client{}
	if len(timeout) > 0 {
		cli.Timeout = timeout[0]
	}
	res, err := cli.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()
	bts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, nil, err
	}
	return res.StatusCode, bts, err
}
func DoHttpObj(method, urls string, body, rets interface{}, timeout ...time.Duration) error {
	if rets == nil {
		return errors.New("rets is nil")
	}
	req, err := NewRequest(method, urls, body)
	if err != nil {
		return err
	}
	cli := &http.Client{}
	if len(timeout) > 0 {
		cli.Timeout = timeout[0]
	}
	res, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	bts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New(string(bts))
	}
	err = json.Unmarshal(bts, rets)
	if err != nil {
		return err
	}
	return nil
}
