package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// function CheckArgs_test checks if checkargs is correct
func TestCheckArgs(t *testing.T) {
	patient := []struct {
		args     []string
		opt      *checkOpt
		expected bool
	}{
		{[]string{"a", "b", "c"}, &CheckOpt{Min: 2, Max: 3}, true},
		{[]string{"a", "b", "c"}, &CheckOpt{Min: 2}, true},
		{[]string{"a", "b", "c"}, &CheckOpt{Min: 2, Max: 2}, false},
		{[]string{"a", "b", "c"}, &CheckOpt{Min: 4}, false},
		{[]string{"a", "b", "c"}, &CheckOpt{Min: 4, Max: 5}, false},
		{[]string{"a", "b", "c"}, &CheckOpt{Max: 2}, false},
		{[]string{"a", "b", "c"}, &CheckOpt{Max: 4}, true},
		{[]string{"a", "b", "c"}, &CheckOpt{Max: 4, Min: 2}, true},
	}

	for _, p := range patient {
		got := CheckArgs(p.args, p.opt)
		assert.Equal(t, p.expected, got, "CheckArgs(%v, %v) = %v, want %v", p.args, p.opt, got, p.expected)
	}
}
