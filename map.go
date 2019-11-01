package ruisUtil

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type Map map[string]interface{}

func NewMap() *Map {
	return &Map{}
}
func NewMaps(body string) *Map {
	e := &Map{}
	json.Unmarshal([]byte(body), e)
	return e
}
func NewMapo(body interface{}) *Map {
	e := &Map{}
	if body == nil {
		return e
	}
	switch body.(type) {
	case string:
		json.Unmarshal([]byte(body.(string)), e)
		break
	case []byte:
		json.Unmarshal(body.([]byte), e)
		break
	case map[string]interface{}:
		for k, v := range body.(map[string]interface{}) {
			e.Set(k, v)
		}
		break
	default:
		bts, err := json.Marshal(body)
		if err == nil {
			json.Unmarshal(bts, e)
		}
		break
	}
	return e
}

func (e *Map) Get(key string) interface{} {
	return (*e)[key]
}
func (e *Map) Set(key string, val interface{}) {
	(*e)[key] = val
}
func (e *Map) Map() map[string]interface{} {
	defer Recovers("convert pointer", nil)
	return *e
}
func (e *Map) ToString() string {
	bts, _ := json.Marshal(e)
	if bts == nil || len(bts) <= 0 {
		return ""
	}
	return string(bts)
}
func (e *Map) GetString(key string) string {
	if e.Get(key) == nil {
		return ""
	}

	return fmt.Sprint(e.Get(key))
}
func (e *Map) GetInt(key string) (int64, error) {
	if e.Get(key) == nil {
		return 0, errors.New("not found")
	}

	v := e.Get(key)
	switch v.(type) {
	case int:
		return v.(int64), nil
	case string:
		return strconv.ParseInt(v.(string), 10, 64)
	case int64:
		return v.(int64), nil
	case float32:
		return int64(v.(float32)), nil
	case float64:
		return int64(v.(float64)), nil
	}
	return 0, errors.New("not found")
}
func (e *Map) GetFloat(key string) (float64, error) {
	if e.Get(key) == nil {
		return 0, errors.New("not found")
	}

	v := e.Get(key)
	switch v.(type) {
	case int:
		return float64(v.(int)), nil
	case string:
		return strconv.ParseFloat(v.(string), 64)
	case int64:
		return float64(v.(int64)), nil
	case float32:
		return float64(v.(float32)), nil
	case float64:
		return v.(float64), nil
	}
	return 0, errors.New("not found")
}
func (e *Map) GetBool(key string) bool {
	if e.Get(key) == nil {
		return false
	}

	v := e.Get(key)
	switch v.(type) {
	case bool:
		return v.(bool)
	}
	return false
}
