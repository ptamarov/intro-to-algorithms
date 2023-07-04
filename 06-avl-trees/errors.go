package avltrees

type OperationErr struct {
	msg string
}

func (oe OperationErr) Error() string {
	return oe.msg
}
