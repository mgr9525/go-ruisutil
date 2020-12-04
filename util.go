package ruisUtil

import (
	"bytes"
	"crypto/md5"
	randc "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync/atomic"
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

// NewXid 生成唯一ID
func NewXid() string {
	return hex.EncodeToString(newXid())
}

// NewXid64 生成唯一ID-base64
func NewXid64() string {
	return base64.StdEncoding.EncodeToString(newXid())
}

func newXid() []byte {
	var b [12]byte
	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(b[:], uint32(time.Now().Unix()))
	// binary.BigEndian.PutUint32(b[:], uint32(time.Now().UnixNano()/1e6))
	// Machine, first 3 bytes of md5(hostname)
	b[4] = machineId[0]
	b[5] = machineId[1]
	b[6] = machineId[2]
	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
	pid := os.Getpid()
	b[7] = byte(pid >> 8)
	b[8] = byte(pid)
	// Increment, 3 bytes, big endian
	i := atomic.AddUint32(&objectIdCounter, 1)
	b[9] = byte(i >> 16)
	b[10] = byte(i >> 8)
	b[11] = byte(i)
	return b[:]
}

// objectIdCounter is atomically incremented when generating a new Xid
var objectIdCounter uint32 = 0

// machineId stores machine id generated once and used in subsequent calls
var machineId = readMachineId()

// initMachineId generates machine id and puts it into the machineId global
// variable. If this function fails to get the hostname, it will cause
// a runtime error.
func readMachineId() []byte {
	var sum [3]byte
	id := sum[:]
	hostname, err1 := os.Hostname()
	if err1 != nil {
		_, err2 := io.ReadFull(randc.Reader, id)
		if err2 != nil {
			panic(fmt.Errorf("cannot get hostname: %v; %v", err1, err2))
		}
		return id
	}
	hw := md5.New()
	hw.Write([]byte(hostname))
	copy(id, hw.Sum(nil))
	return id
}
