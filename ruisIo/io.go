package ruisIo

import (
	"context"
	"errors"
	"fmt"
	"io"
)

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
			//return nil, fmt.Errorf("read ln=0 broker")
			break
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
			//return fmt.Errorf("write ln=0 broker")
			break
		}
	}
	if wn != ln {
		return fmt.Errorf("write len err:%d/%d", wn, ln)
	}
	return nil
}
