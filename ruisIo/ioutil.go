package ruisIo

import (
	"context"
	"errors"
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func CheckContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return errors.New("end")
	default:
		return nil
	}
}
