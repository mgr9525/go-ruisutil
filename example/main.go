package main

import (
	"encoding/hex"
	"fmt"
	"github.com/mgr9525/go-ruisutil/ruisCgo"
	"unsafe"
)

type flvHeader struct {
	Sig  [3]byte
	Ver  byte
	Flag byte
	Len  int32
}

func main() {
	hdr := &flvHeader{}
	hds, err := hex.DecodeString("464c560105000000090000000012000111000000")
	if err != nil {
		println(err.Error())
		return
	}
	var i1 int32 = 1
	println("sizeof1:", unsafe.Sizeof(i1))
	println("sizeof:", unsafe.Sizeof(*hdr), "alif:", unsafe.Alignof(*hdr))
	err = ruisCgo.Bytes2Struct(unsafe.Pointer(hdr), hds, unsafe.Sizeof(*hdr))
	//err=ruisCgo.Bytes2Structs(unsafe.Pointer(hdr),hds)
	println(fmt.Sprintf("sig1:%x,hdr.Ver:%x,.Flag:%x,ln:%d,err=%+v", hdr.Sig[1], hdr.Ver, hdr.Flag, hdr.Len, err))
	bts, err := ruisCgo.Struct2Bytes(unsafe.Pointer(hdr), unsafe.Sizeof(*hdr))
	if err != nil {
		println(err.Error())
		return
	}
	println(hex.EncodeToString(bts))
}
