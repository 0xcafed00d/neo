package neo

import (
	"fmt"
	"os"
	"strings"
)

func ExitOnError(e error, exitcode int) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(exitcode)
	}
}

func PanicOnError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		panic(e)
	}
}

// same as strings.TrimPrefix, but also returns a bool to indicate if the prefix was trimmed
func TryTrimPrefix(s, prefix string) (string, bool) {
	has := strings.HasPrefix(s, prefix)
	if has {
		return s[len(prefix):], true
	}
	return s, false
}

type ErrorStr string

func (e ErrorStr) Error() string { return string(e) }

type ErrorWrapper struct {
	Message string
	Err     error
}

func (e *ErrorWrapper) Error() string { return e.Message + ": " + e.Err.Error() }

type Named interface {
	Name() string
}

type Nameable interface {
	Named
	SetName(name string)
}
