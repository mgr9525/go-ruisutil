package ruisIo

import (
	"context"
	"errors"
	"io"
)

func TcpRead(ctx context.Context, conn io.Reader, ln uint) ([]byte, error) {
	if conn == nil || ln <= 0 {
		return nil, errors.New("handleRead ln<0")
	}
	rn := uint(0)
	bts := make([]byte, ln)
	for {
		if EndContext(ctx) {
			return nil, errors.New("context dead")
		}
		n, err := conn.Read(bts[rn:])
		if n > 0 {
			rn += uint(n)
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
	return bts, nil
}
func TcpWrite(ctx context.Context, conn io.Writer, bts []byte) error {
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
