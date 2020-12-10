package ruisUtil

import (
	"context"
	"errors"
	"io"
	"time"
)

type CircleByteBuffer struct {
	io.Reader
	io.Writer
	io.Closer
	datas []byte

	start int
	end   int
	size  int

	ctx  context.Context
	cncl context.CancelFunc
}

func NewCircleByteBuffer(ctx context.Context, len int) *CircleByteBuffer {
	var e = new(CircleByteBuffer)
	e.datas = make([]byte, len)
	e.start = 0
	e.end = 0
	e.size = len
	e.ctx, e.cncl = context.WithCancel(ctx)
	return e
}

func (e *CircleByteBuffer) GetLen() int {
	if e.start == e.end {
		return 0
	} else if e.start < e.end {
		return e.end - e.start
	} else {
		return e.size - (e.start - e.end)
	}
}
func (e *CircleByteBuffer) GetFree() int {
	return e.size - e.GetLen()
}
func (e *CircleByteBuffer) Clear() {
	e.start = 0
	e.end = 0
}
func (e *CircleByteBuffer) PutByte(b byte) error {
	if e.IsClose() {
		return io.EOF
	}
	e.datas[e.end] = b
	var pos = e.end + 1
	if pos == e.size {
		pos = 0
	}
	for pos == e.start {
		if e.IsClose() {
			return io.EOF
		}
		time.Sleep(time.Millisecond)
	}
	e.end = pos
	return nil
}

func (e *CircleByteBuffer) GetByte() (byte, error) {
	if e.IsClose() {
		return 0, io.EOF
	}
	for e.GetLen() <= 0 {
		if e.IsClose() {
			return 0, io.EOF
		}
		time.Sleep(time.Millisecond)
	}
	var ret = e.datas[e.start]
	pos := e.start + 1
	if pos == e.size {
		pos = 0
	}
	e.start = pos
	return ret, nil
}
func (e *CircleByteBuffer) Geti(i int) byte {
	if i >= e.GetLen() {
		panic("out buffer")
	}
	var pos = e.start + i
	if pos >= e.size {
		pos -= e.size
	}
	return e.datas[pos]
}

/*func (e*CircleByteBuffer)puts(bts []byte){
	for i:=0;i<len(bts);i++{
		e.put(bts[i])
	}
}
func (e*CircleByteBuffer)gets(bts []byte)int{
	if bts==nil {return 0}
	var ret=0
	for i:=0;i<len(bts);i++{
		if e.GetLen()<=0{break}
		bts[i]=e.get()
		ret++
	}
	return ret
}*/
func (e *CircleByteBuffer) IsClose() bool {
	return CheckCtxDone(e.ctx)
}
func (e *CircleByteBuffer) Close() error {
	if e.IsClose() || e.cncl == nil {
		return errors.New("already closed")
	}
	e.cncl()
	e.cncl = nil
	return nil
}
func (e *CircleByteBuffer) Read(bts []byte) (int, error) {
	if e.IsClose() {
		return 0, io.EOF
	}
	if bts == nil {
		return 0, errors.New("bts is nil")
	}
	var ret = 0
	for i := 0; i < len(bts); i++ {
		b, err := e.GetByte()
		if err != nil {
			return ret, err
		}
		bts[i] = b
		ret++
	}
	return ret, nil
}
func (e *CircleByteBuffer) Write(bts []byte) (int, error) {
	if e.IsClose() {
		return 0, io.EOF
	}
	var ret = 0
	for i := 0; i < len(bts); i++ {
		err := e.PutByte(bts[i])
		if err != nil {
			return ret, err
		}
		ret++
	}
	return ret, nil
}
