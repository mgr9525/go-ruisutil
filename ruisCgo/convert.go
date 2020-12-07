package ruisCgo

// #include<string.h>
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
)

type SliceMock struct {
	addr unsafe.Pointer
	len  int
	cap  int
}

func Struct2Bytes(pt interface{}) ([]byte, error) {
	if pt == nil {
		return nil, errors.New("param is nil")
	}
	vf := reflect.ValueOf(pt)
	if vf.IsNil() {
		return nil, errors.New("param is err")
	}
	data := pt
	if vf.Type().Kind() == reflect.Ptr {
		data = vf.Elem().Interface()
	}
	ln := unsafe.Sizeof(data)
	mock := &SliceMock{
		addr: unsafe.Pointer(&data),
		cap:  int(ln),
		len:  int(ln),
	}
	bts := *(*[]byte)(unsafe.Pointer(mock))
	rtbts := make([]byte, mock.len)
	copy(rtbts, bts)
	return rtbts, nil
}
func Bytes2Struct(pt interface{}, bts []byte) error {
	if pt == nil {
		return errors.New("param is nil")
	}
	vf := reflect.ValueOf(pt)
	if vf.Type().Kind() != reflect.Ptr || vf.IsNil() {
		return errors.New("param is err")
	}
	data := vf.Elem().Interface()
	ln := unsafe.Sizeof(data)
	dataptr := uintptr(unsafe.Pointer(&data))
	C.memcpy(unsafe.Pointer(dataptr), unsafe.Pointer(&bts[0]), C.ulong(ln))
	vf.Elem().Set(reflect.ValueOf(data))
	return nil
}
