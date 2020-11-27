package ruisUtil

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"math/rand"
	"time"
)

func Recovers(handle func(errs interface{})) {
	if err := recover(); err != nil {
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

func Int2Bytes(v int64, len int) ([]byte, error) {
	var tmp interface{}
	buf := &bytes.Buffer{}
	switch len {
	case 8:
		t := int8(v)
		tmp = &t
	case 16:
		t := int16(v)
		tmp = &t
	case 32:
		t := int32(v)
		tmp = &t
	case 64:
		tmp = &v
	}
	err := binary.Write(buf, binary.BigEndian, tmp)
	return buf.Bytes(), err
}
func Bytes2Int(v []byte, len int) (ret int64, rterr error) {
	ret = int64(0)
	buf := bytes.NewBuffer(v)
	switch len {
	case 8:
		var tmp int8
		rterr = binary.Read(buf, binary.BigEndian, &tmp)
		ret = int64(tmp)
	case 16:
		var tmp int16
		rterr = binary.Read(buf, binary.BigEndian, &tmp)
		ret = int64(tmp)
	case 32:
		var tmp int32
		rterr = binary.Read(buf, binary.BigEndian, &tmp)
		ret = int64(tmp)
	case 64:
		rterr = binary.Read(buf, binary.BigEndian, &ret)
	}
	return
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

func StructInterturn(src, dist interface{}) error {
	bts, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bts, dist)
	if err != nil {
		return err
	}
	return nil
}

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

// RandomString 随机生成字符串
func RandomString(l int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	bts := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bts[r.Intn(len(bts))])
	}
	return string(result)
}
