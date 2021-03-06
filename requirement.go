// Package requirement provides support for requirement to be able to skip some tests depending on the environment.
// It is intended to work at least with the built-in ``testing'' framework. But any testing framework that
// defines a ``Skip(args ...interface{})'' method can be used.
//
// To write a test with a requirement, you need import this library and call it like the following
//     func TestWithRequirement(t *testing.T) {
//         requirement.Is(t, func() bool {
//             return false
//         })
//     }
//
// This will print
//     requirement.go:51: requirement_test.go:56: unmatched requirement TestWithRequirement.func1
package requirement

import (
	"fmt"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// skipT defines what should be provides by the testing framework
type skipT interface {
	Skip(args ...interface{})
}

// test represent a function that can be used as a requirement validation.
type test func() bool

// Is checks if the environment satisfies the requirements
// for the test to run or skips the tests.
func Is(s skipT, requirements ...test) {
	for _, r := range requirements {
		isValid := r()
		if !isValid {
			requirementFunc := runtime.FuncForPC(reflect.ValueOf(r).Pointer()).Name()
			skip(s, requirementFunc)
		}
	}
}

func skip(s skipT, reason string) {
	var source string
	_, filename, line, ok := runtime.Caller(2)
	if ok {
		source = fmt.Sprintf("%s:%d: ", filepath.Base(filename), line)
	}
	s.Skip(fmt.Sprintf("%sunmatched requirement %s", source, extractRequirement(reason)))
}

func extractRequirement(requirementFunc string) string {
	requirement := path.Base(requirementFunc)
	return strings.SplitN(requirement, ".", 2)[1]
}
