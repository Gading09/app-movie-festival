package response

import "fmt"

func New(statusCode int, msg string, err error) error {
	return fmt.Errorf("%d | %s | %w", statusCode, msg, err)
}
