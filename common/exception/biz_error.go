package exception

import "fmt"

type BizError struct {
	Code    int
	Message string
}

func (e *BizError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
