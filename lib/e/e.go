package e

import "fmt"

func Warp(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func WarpIsErr(msg string, err error) error {
	if err == nil {
		return nil
	}
	return Warp(msg, err)
}
