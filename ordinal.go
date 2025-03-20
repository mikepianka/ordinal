package ordinal

import (
	"strconv"
	"strings"
)

func Itoa(i int) string {
	if i < 0 {
		return "negative"
	}

	// avoid division by zero
	if i == 0 {
		return "0th"
	}

	s := strconv.Itoa(i)

	// account for teens
	if strings.HasSuffix(s, "11") || strings.HasSuffix(s, "12") || strings.HasSuffix(s, "13") {
		return s + "th"
	}

	// any other numbers
	switch i % 10 {
	case 1:
		return s + "st"
	case 2:
		return s + "nd"
	case 3:
		return s + "rd"
	default:
		return s + "th"
	}
}
