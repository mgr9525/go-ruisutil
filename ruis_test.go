package ruisUtil

import (
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	m := NewMap()
	m.Set("tm", time.Now())
	m.Set("tit", "hahaha!!")
	println(m.ToString())
	m.Set("zh", "ok")
	mt := NewMaps(m.ToString())
	println(mt.ToString())
}
