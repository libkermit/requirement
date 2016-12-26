package requirement

import (
	"testing"
)

type customSkip struct {
	reason string
}

func (s *customSkip) Skip(args ...interface{}) {
	s.reason = args[0].(string)
}

func alwaysFalse() bool {
	return false
}

func alwaysTrue() bool {
	return true
}

func TestIsFuncTrue(t *testing.T) {
	s := &customSkip{}
	Is(s, alwaysTrue)
	expected := ""
	if s.reason != expected {
		t.Fatalf("expected reason %q, got %q", expected, s.reason)
	}
}

func TestIsFuncFalse(t *testing.T) {
	s := &customSkip{}
	Is(s, alwaysFalse)
	expected := "unmatched requirement alwaysFalse"
	if s.reason != expected {
		t.Fatalf("expected reason %q, got %q", expected, s.reason)
	}
}

func TestIsAnonymousFunc(t *testing.T) {
	s := &customSkip{}
	Is(s, func() bool {
		return false
	})
	expected := "unmatched requirement TestIsAnonymousFunc.func1"
	if s.reason != expected {
		t.Fatalf("expected reason %q, got %q", expected, s.reason)
	}
}
