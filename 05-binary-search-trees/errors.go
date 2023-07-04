package bsts

import "fmt"

type OperationErr struct {
	msg string
}

func (oe OperationErr) Error() string {
	return fmt.Sprint(oe.msg)
}
