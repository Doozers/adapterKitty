package utils

import (
	"fmt"
)

type CheckOpt struct {
	Min int
	Max int
}

func CheckArgs(args []string, opt *CheckOpt) bool {
	if opt == nil {
		return true
	}

	if len(args) < opt.Min {
		fmt.Printf("not enough arguments: expected %d, got %d\n", opt.Min, len(args))
		return false
	}

	if opt.Max == 0 && opt.Min != 0 {
		return true
	}

	if len(args) > opt.Max {
		fmt.Printf("too many arguments: expected %d, got %d\n", opt.Max, len(args))
		return false
	}

	return true
}
