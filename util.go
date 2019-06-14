package ruisUtil

import "fmt"

func Recovers(name string) {
	if err := recover(); err != nil {
		fmt.Print("ruisRecover(" + name + "):")
		fmt.Println(err) // 这里的err其实就是panic传入的内容，55
	}
}

func Ruis_bytes2short(bts []byte) int {
	var ret = 0
	if len(bts) >= 2 {
		//fmt.Println("ruis_bytes2short:",(int(bts[0])<<8)&0xff),"，",(int(bts[0])&0xff))
		ret = (int(int(bts[0])<<8) & 0xffff) | (int(bts[1]) & 0xff)
	}
	return ret
}
func Ruis_short2bytes(v int) []byte {
	var bts = []byte{0, 0}
	bts[0] = byte((v >> 8) & 0xff)
	bts[1] = byte(v & 0xff)
	return bts
}
