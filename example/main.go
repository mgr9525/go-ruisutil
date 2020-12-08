package main

import (
	"encoding/hex"
	"fmt"
	"github.com/mgr9525/go-ruisutil/ruisCgo"
	"unsafe"
)

type FlvHeader struct {
	Sig  [3]byte
	Ver  byte
	Flag byte
	Len  int32
}

func main() {
	hdr := &FlvHeader{}
	hds, err := hex.DecodeString("464c560105000000090000000012000111000000")
	if err != nil {
		println(err.Error())
		return
	}
	err = ruisCgo.Bytes2Struct(unsafe.Pointer(hdr), hds, unsafe.Sizeof(*hdr))
	//err=ruisCgo.Bytes2Structs(unsafe.Pointer(hdr),hds)
	println(fmt.Sprintf("sig1:%x,hdr.Ver:%x,.Flag:%x,err=%+v", hdr.Sig[1], hdr.Ver, hdr.Flag, err))
	bts, err := ruisCgo.Struct2Bytes(unsafe.Pointer(hdr), unsafe.Sizeof(*hdr))
	if err != nil {
		println(err.Error())
		return
	}
	println(hex.EncodeToString(bts))
}
