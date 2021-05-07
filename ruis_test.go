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
func Test2(t *testing.T) {
	println(fmt.Sprint(float64(554646954654.00)))
}
