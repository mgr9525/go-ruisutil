package ruisIo

import (
	"context"
	"errors"
	"fmt"
	"io"
)

/*func TcpRead(ctx context.Context, conn io.Reader, ln uint) ([]byte, error) {
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
}*/

/*func TcpWrite(ctx context.Context, conn io.Writer, bts []byte) error {
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
		wn += n
		if err != nil {
			return err
		}
		if n <= 0 && wn < ln {
			return errors.New("conn abort")
		}
	}
	return nil
}*/

func IoReadAll(ctx context.Context, rdr io.Reader, ln uint) ([]byte, error) {
	rn := uint(0)
	bts := make([]byte, ln)
	for {
		if EndContext(ctx) {
			return nil, errors.New("ctx end")
		}
		n, err := rdr.Read(bts[rn:])
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
			return nil, fmt.Errorf("read ln=0 broker")
		}
	}
	if rn != ln {
		return nil, fmt.Errorf("read len err:%d/%d", rn, ln)
	}
	return bts, nil
}
func IoWriteAll(ctx context.Context, wtr io.Writer, bts []byte) error {
	wn := 0
	ln := len(bts)
	for {
		if EndContext(ctx) {
			return errors.New("ctx end")
		}
		n, err := wtr.Write(bts[wn:])
		if n > 0 {
			wn += n
		}
		if wn >= ln {
			break
		}
		if err != nil {
			return err
		}
		if n <= 0 {
			return fmt.Errorf("write ln=0 broker")
		}
	}
	if wn != ln {
		return fmt.Errorf("write len err:%d/%d", wn, ln)
	}
	return nil
}
