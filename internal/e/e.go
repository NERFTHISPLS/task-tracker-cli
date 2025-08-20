package e

import "fmt"

func Wrap(msg string, e error) error {
	return fmt.Errorf("%s: %w", msg, e)
}
