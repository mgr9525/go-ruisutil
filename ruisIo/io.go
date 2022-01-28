package ruisIo

import (
	"bytes"
	"context"
	"errors"
	"net"
)

func TcpRead(ctx context.Context, conn net.Conn, ln uint) ([]byte, error) {
	if conn == nil || ln <= 0 {
		return nil, errors.New("handleRead ln<0")
	}
	var buf *bytes.Buffer
	rn := uint(0)
	tn := ln
	if ln > 10240 {
		tn = 10240
		buf = &bytes.Buffer{}
	}
	bts := make([]byte, tn)
	for {
		if EndContext(ctx) {
			return nil, errors.New("context dead")
		}
		n, err := conn.Read(bts)
		if n > 0 {
			rn += uint(n)
			if buf != nil {
				buf.Write(bts[:n])
			}
		}
		if rn >= ln {
			break
		}
		if err != nil {
			return nil, err
		}
		if n <= 0 {
			return nil, errors.New("conn abort")
		}
	}
	if buf != nil {
		return buf.Bytes(), nil
	}
	return bts, nil
}
func TcpWrite(ctx context.Context, conn net.Conn, bts []byte) error {
	ln := len(bts)
	if conn == nil || ln <= 0 {
		return errors.New("handleRead ln<0")
	}

	wn := 0
	for wn < ln {
		if EndContext(ctx) {
			return errors.New("context dead")
		}
		n, err := conn.Write(bts[wn:])
		if err != nil {
			return err
		}
		wn += n
	}
	return nil
}
