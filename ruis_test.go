package ruisUtil

import (
	"fmt"
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

	bts, _ := Int2Bytes(123, 16)
	println(fmt.Sprintf("Int2Bytes:%x", bts))
	n, _ := Bytes2Int(bts, 32)
	println(fmt.Sprintf("Bytes2Int:%d", n))
}

func add(z *int) int {
	println("add200:", *z)
	*z += 200
	return *z
}
func deferRet(x, y int) (z int) {
	defer func() {
		println("info:", z)
		z += 100
	}()
	z = x + y
	return add(&z)
}
func Test1(t *testing.T) {
	i := deferRet(1, 1)
	println("main:", i)
}
