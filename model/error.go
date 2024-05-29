package model

import "fmt"

type (
	CutomError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func (e CutomError) Error() error {
	return fmt.Errorf("code: %d, message: %s", e.Code, e.Message)
}
