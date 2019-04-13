package lifecycle

import "strings"

type multiError struct {
	errChan   chan error
	errors    []string
	separator string
}

// Error implements error interface
func (e *multiError) Error() string {
	return strings.Join(e.errors, e.separator)
}

func (e *multiError) collect() {
	for ee := range e.errChan {
		if ee != nil {
			e.errors = append(e.errors, ee.Error())
		}
	}
}
func (e *multiError) ok() bool {
	return len(e.errors) == 0
}
