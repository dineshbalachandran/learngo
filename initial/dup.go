package main

import (
	"bufio"
	"fmt"
	"os"
)

func dup() {
	counts := make(map[string]int)
	//map[string]struct{} is the way to get a "set" of strings in Go!!
	dupfilenames := make(map[string]map[string]struct{})

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, dupfilenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts, dupfilenames)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("%s", "Files:")
			for name := range dupfilenames[line] {
				fmt.Printf("\t%s\n", name)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, dupfilenames map[string]map[string]struct{}) {

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if dupfilenames[input.Text()] == nil {
			dupfilenames[input.Text()] = make(map[string]struct{})
		}
		dupfilenames[input.Text()][f.Name()] = struct{}{}
	}
}
