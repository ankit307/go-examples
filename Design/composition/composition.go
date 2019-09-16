package main

import (
	"fmt"
	"strings"
)

type fnString func(string) string

func composeRec(fns ...fnString) fnString {
	return func(s string) string {
		f := fns[0]
		fs := fns[1:len(fns)]

		if len(fns) == 1 {
			return f(s)
		}

		return f(composeRec(fs...)(s))
	}
}

func main() {
	raw := "\n\n\nHello Golang!!!\n\n\n"
	trim := strings.TrimSpace
	trimExclamation := func(s string) string {
		return strings.Trim(s, "!")
	}
	toLower := strings.ToLower
	format := composeRec(toLower, trimExclamation, trim)

	fmt.Println(format(raw))
}
