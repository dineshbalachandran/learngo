package main

import "strings"

func stringsjoin(ss ...string) string {
	var as []string
	for _, s := range ss {
		as = append(as, s)
	}

	return strings.Join(as, " ")
}
