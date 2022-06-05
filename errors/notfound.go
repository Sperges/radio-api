package errors

import (
	"fmt"
	"net/http"
)

type NotFoundError struct {
	err string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf(e.err)
}

func (e *NotFoundError) Code() int {
	return http.StatusNotFound
}
