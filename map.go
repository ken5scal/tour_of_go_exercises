package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		(ret[words[i]])++
	}

	return ret
}

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	wc.Test(WordCount)
}
