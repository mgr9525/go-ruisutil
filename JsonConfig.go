package ruisUtil

import (
	"encoding/json"
	"fmt"
	"github.com/mgr9525/go-ruisutil/ruisio"
	"io/ioutil"
)

type JsonConfig struct {
	path string
	root map[string]interface{}
}

func NewJsonConfig(path string) *JsonConfig {
	var e = new(JsonConfig)
	e.path = path
	e.root = make(map[string]interface{})
	e.init()
	return e
}

func (e *JsonConfig) init() {
	defer Recovers("JsonConfig.init", nil)
	if e.path != "" && ruisIo.PathExists(e.path) {
		conts, err := ioutil.ReadFile(e.path)
		if err == nil {
			fmt.Println("readTo:" + e.path + "! conts:" + string(conts))
			err := json.Unmarshal(conts, &e.root)
			if err != nil {
				fmt.Println("readTo:err")
				e.root = map[string]interface{}{}
			}
		}
	}
	fmt.Println("init:ok")
}

func (e *JsonConfig) Reinit(bts []byte) {
	defer Recovers("JsonConfig.reinit", nil)
	json.Unmarshal(bts, &e.root)
}

func (e *JsonConfig) Gets(key string) string {
	ret, ext := e.root[key]
	if ext {
		return ret.(string)
	}
	return ""
}
func (e *JsonConfig) Sets(key string, value string) {
	e.root[key] = value
	go e.Save()
}

func (e *JsonConfig) Save() {
	defer Recovers("JsonConfig.save", nil)

	if e.path != "" {
		js, err := json.Marshal(e.root)
		if err == nil {
			//fmt.Println("saveTo:"+ruisApp.dataPath+e.name+"! js:"+string(js))
			err = ioutil.WriteFile(e.path, js, 0644)
			if err != nil {
				fmt.Println("conf save err")
			}
		} else {
			fmt.Println("conf save err")
		}
	}
}
