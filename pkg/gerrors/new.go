package gerrors

import (
	"errors"
	"strings"
	"sync"
)

var _ IError = (*Error)(nil)

type (
	IError interface {
		error
		WithFields(fields map[string]string) IError
		WithMessage(message string) IError
		WithCode(code Code) IError
		WithError(err error) IError
		AddFields(field, msg string) IError
		IsNil() error
	}

	Error struct {
		message []string
		Code    Code              `json:"code"`
		Msg     string            `json:"msg"`
		Fields  map[string]string `json:"fields"`
	}

	ErrorOption func(*Error)
)

var lock sync.Mutex

// AddFields implements IError.
func (l *Error) AddFields(field string, msg string) IError {
	lock.Lock()
	defer lock.Unlock()
	l.Fields[field] = msg
	l.setMsg()
	return l
}

// WithCode implements IError.
func (l *Error) WithCode(code Code) IError {
	lock.Lock()
	defer lock.Unlock()
	l.Code = code
	l.setMsg()
	return l
}

// WithFields implements IError.
func (l *Error) WithFields(fields map[string]string) IError {
	lock.Lock()
	defer lock.Unlock()
	l.Fields = fields
	l.setMsg()
	return l
}

// WithMessage implements IError.
func (l *Error) WithMessage(message string) IError {
	lock.Lock()
	defer lock.Unlock()
	l.message = append(l.message, message)
	l.setMsg()
	return l
}

// WithError implements IError.
func (l *Error) WithError(err error) IError {
	if err == nil {
		return l
	}
	lock.Lock()
	defer lock.Unlock()
	var e *Error
	if errors.As(err, &e) {
		l.Code = e.Code
		l.message = append(l.message, e.message...)
		for k, v := range e.Fields {
			l.Fields[k] = v
		}
	} else {
		l.message = append(l.message, err.Error())
	}
	l.setMsg()

	return l
}

// IsNil implements IError.
func (l *Error) IsNil() error {
	if (l.message == nil || len(l.message) == 0) && (l.Fields == nil || len(l.Fields) == 0) && l.Code == 0 {
		return nil
	}
	l.setMsg()
	return l
}

// setMsg implements IError.
func (l *Error) setMsg() {
	if len(l.message) > 0 {
		l.Msg = l.message[len(l.message)-1]
		return
	}
	l.Msg = l.Error()
}

func (l *Error) Error() string {
	switch len(l.message) {
	case 0:
		if l.Code == 0 {
			return ""
		}
		return l.Code.Error().Error()
	case 1:
		return l.message[0]
	default:
		buff := strings.Builder{}
		// "(a) --> (b) --> (c)"
		for i := len(l.message) - 1; i >= 0; i-- {
			buff.WriteString("(")
			buff.WriteString(l.message[i])
			buff.WriteString(")")
			if i != 0 {
				buff.WriteString(" --> ")
			}
		}

		return buff.String()
	}
}

func New(opts ...ErrorOption) *Error {
	e := &Error{
		Code:   Unknown,
		Fields: make(map[string]string),
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

func NewNil() *Error {
	return &Error{
		Fields: make(map[string]string),
	}
}

func WithFields(fields map[string]string) ErrorOption {
	return func(e *Error) {
		e.Fields = fields
	}
}

func WithMessage(message string) ErrorOption {
	return func(e *Error) {
		e.message = append(e.message, message)
	}
}

func WithCode(code Code) ErrorOption {
	return func(e *Error) {
		e.Code = code
	}
}

func Is(err error, code Code) bool {
	var e *Error
	if errors.As(err, &e) {
		return e.Code == code
	}
	return false
}

func Equal(err1, err2 error) bool {
	var e1 *Error
	if errors.As(err1, &e1) {
		var e2 *Error
		if errors.As(err2, &e2) {
			return e1.Code == e2.Code
		}
	}
	return false
}
