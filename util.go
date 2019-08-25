package ruisUtil

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type ErrHandle func(o interface{})

func Recovers(name string, handle ErrHandle) {
	if err := recover(); err != nil {
		fmt.Print("ruisRecover(" + name + "):")
		fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		if handle != nil {
			handle(err)
		}
	}
}

func Bytes2Short(bts []byte) int {
	var ret = 0
	if len(bts) >= 2 {
		ret = (int(int(bts[0])<<8) & 0xffff) | (int(bts[1]) & 0xff)
	}
	return ret
}
func Short2Bytes(v int) []byte {
	var bts = []byte{0, 0}
	bts[0] = byte((v >> 8) & 0xff)
	bts[1] = byte(v & 0xff)
	return bts
}

func Int2Bytes(v int64, len int) []byte {
	buf := &bytes.Buffer{}
	switch len {
	case 8:
		var tmp = int8(v)
		err := binary.Write(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Write err")
		}
		break
	case 16:
		var tmp = int16(v)
		err := binary.Write(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Write err")
		}
		break
	case 32:
		var tmp = int32(v)
		err := binary.Write(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Write err")
		}
		break
	case 64:
		var tmp = int64(v)
		err := binary.Write(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Write err")
		}
		break
	}
	return buf.Bytes()
}
func Bytes2Int(v []byte, len int) int64 {
	buf := bytes.NewBuffer(v)
	switch len {
	case 8:
		var tmp int8
		err := binary.Read(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Read err")
		}
		return int64(tmp)
	case 16:
		var tmp int16
		err := binary.Read(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Read err")
		}
		return int64(tmp)
	case 32:
		var tmp int32
		err := binary.Read(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Read err")
		}
		return int64(tmp)
	case 64:
		var tmp int64
		err := binary.Read(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Read err")
		}
		return int64(tmp)
	}

	return 0
}
func UInt2Bytes(v int64, len int) []byte {
	buf := &bytes.Buffer{}
	switch len {
	case 8:
		var tmp = uint8(v)
		err := binary.Write(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Write err")
		}
		break
	case 16:
		var tmp = uint16(v)
		err := binary.Write(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Write err:" + err.Error())
		}
		break
	case 32:
		var tmp = uint32(v)
		err := binary.Write(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Write err:" + err.Error())
		}
		break
	case 64:
		var tmp = uint64(v)
		err := binary.Write(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Write err")
		}
		break
	}
	return buf.Bytes()
}
func Bytes2UInt(v []byte, len int) uint64 {
	buf := bytes.NewBuffer(v)
	switch len {
	case 8:
		var tmp uint8
		err := binary.Read(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Read err")
		}
		return uint64(tmp)
	case 16:
		var tmp uint16
		err := binary.Read(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Read err")
		}
		return uint64(tmp)
	case 32:
		var tmp uint32
		err := binary.Read(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Read err")
		}
		return uint64(tmp)
	case 64:
		var tmp uint64
		err := binary.Read(buf, binary.BigEndian, &tmp)
		if err != nil {
			panic("binary.Read err")
		}
		return uint64(tmp)
	}

	return 0
}
