package messenger

import (
	"strings"
	"testing"
)

// TestLogLevel verifies that we do not get colored output if we disable it.
func TestLogLevel(t *testing.T) {
	log := New(NOTSET, false)

	var b strings.Builder
	log.out = &b
	log.Error("No Color")

	given := b.String()
	expected := quote(log, "No Color")

	if given != expected {
		t.Errorf("Given %s, Expected %s", given, expected)
	}
}

// quote appends the terminal reset and newline to a given string.
func quote(m *messenger, msg string) string {
	return msg + string(m.reset()) + "\n"
}

func unquote(m *messenger, msg string) string {
	s := strings.TrimSuffix(msg, "\n")
	return strings.TrimSuffix(s, string(Reset))
}
