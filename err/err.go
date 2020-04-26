package err

import (
	"log"
)

type ErrorCode = int

type IErr interface {
	Error() string
	ErrCode() int
	Log()
	HttpStatus() int
}

type Err struct {
	Code ErrorCode
	Msg string
	Status int
}

func (e *Err) Error() string {
	return e.Msg
}

func (e *Err) Log() {
	log.Println(e.Msg)
}

func (e *Err) HttpStatus() int {
	return e.Status
}

func (e *Err) ErrCode() int {
	return e.Code
}