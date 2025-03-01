package gofp

import (
	"fmt"

	"golang.org/x/xerrors"
)

type XError struct {
	Err       error
	Message   string
	Stack     xerrors.Frame
	ErrorCode int
}

func (m *XError) Error() string {
	return m.Message
}

func (m *XError) Format(f fmt.State, c rune) { // implements fmt.Formatter
	xerrors.FormatError(m, f, c)
}

const (
	newlineTab = "\n\t"
	one        = 1
)

func (m *XError) FormatError(p xerrors.Printer) error { // implements xerrors.Formatter
	p.Print(m.Message)
	p.Print(newlineTab)
	p.Print("Error code: " + fmt.Sprint(m.ErrorCode))
	p.Print(newlineTab)
	p.Print(m.Err)
	if p.Detail() {
		m.Stack.Format(p)
	}
	return nil
}

const defaultStackDepth = 1

func NewXError(message string, err error, errCode int) XError {
	return XError{
		Message:   message,
		Err:       err,
		Stack:     xerrors.Caller(defaultStackDepth),
		ErrorCode: errCode,
	}
}

func NewXErrorF(message string, errCode int) XError {
	return XError{
		Message:   message,
		Err:       fmt.Errorf(message, nil),
		Stack:     xerrors.Caller(one),
		ErrorCode: errCode,
	}
}
