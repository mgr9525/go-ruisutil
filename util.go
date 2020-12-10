package ruisUtil

import (
	"bytes"
	"context"
	"crypto/md5"
	randc "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math"
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

// BigEndian
func BigByteToInt(data []byte) int64 {
	ln := len(data)
	rt := int64(0)
	for i := 0; i < ln; i++ {
		rt |= int64(data[ln-1-i]) << (i * 8)
	}
	return rt
}
func BigIntToByte(data int64, ln int) []byte {
	rt := make([]byte, ln)
	for i := 0; i < ln; i++ {
		rt[ln-1-i] = byte(data >> (i * 8))
	}
	return rt
}

// LittleEndian
func LitByteToInt(data []byte) int64 {
	ln := len(data)
	rt := int64(0)
	for i := 0; i < ln; i++ {
		rt |= int64(data[i]) << (i * 8)
	}
	return rt
}
func LitIntToByte(data int64, ln int) []byte {
	rt := make([]byte, ln)
	for i := 0; i < ln; i++ {
		rt[i] = byte(data >> (i * 8))
	}
	return rt
}

// BigEndian
func BigByteToFloat32(data []byte) float32 {
	v := BigByteToInt(data)
	return math.Float32frombits(uint32(v))
}
func BigFloatToByte32(data float32) []byte {
	v := math.Float32bits(data)
	return BigIntToByte(int64(v), 4)
}
func BigByteToFloat64(data []byte) float64 {
	v := BigByteToInt(data)
	return math.Float64frombits(uint64(v))
}
func BigFloatToByte64(data float64) []byte {
	v := math.Float64bits(data)
	return BigIntToByte(int64(v), 8)
}

// LittleEndian
func LitByteToFloat32(data []byte) float32 {
	v := LitByteToInt(data)
	return math.Float32frombits(uint32(v))
}
func LitFloatToByte32(data float32) []byte {
	v := math.Float32bits(data)
	return LitIntToByte(int64(v), 4)
}
func LitByteToFloat64(data []byte) float64 {
	v := LitByteToInt(data)
	return math.Float64frombits(uint64(v))
}
func LitFloatToByte64(data float64) []byte {
	v := math.Float64bits(data)
	return LitIntToByte(int64(v), 8)
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
	return base64.RawURLEncoding.EncodeToString(newXid())
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

func CheckCtxDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
