package ruisCgo

/* #include<string.h>
void ruismemcpy(void *dst,void*src,unsigned long ln){
	memcpy(dst,src,(size_t)ln);
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

type SliceMock struct {
	addr unsafe.Pointer
	len  int
	cap  int
}

/*func Struct2Bytes(pt interface{}) ([]byte, error) {
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
	C.memcpy(unsafe.Pointer(dataptr), unsafe.Pointer(&bts[0]), C.ulonglong(ln))
	vf.Elem().Set(reflect.ValueOf(data))
	return nil
}*/
func Struct2Bytes(pt unsafe.Pointer, ln uint) ([]byte, error) {
	mock := &SliceMock{
		addr: pt,
		cap:  int(ln),
		len:  int(ln),
	}
	bts := *(*[]byte)(unsafe.Pointer(mock))
	rtbts := make([]byte, mock.len)
	copy(rtbts, bts)
	return rtbts, nil
}
func Bytes2Struct(pt unsafe.Pointer, bts []byte, ln uint) error {
	if pt == nil {
		return errors.New("param is nil")
	}
	dataptr := uintptr(pt)
	C.ruismemcpy(unsafe.Pointer(dataptr), unsafe.Pointer(&bts[0]), C.ulong(ln))
	return nil
}
