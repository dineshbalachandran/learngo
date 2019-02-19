package main

import (
	"fmt"
	"sort"
)

func topoSort(m map[string][]string) []string {
	var order []string
	set := make(map[string]bool)
	seen := make(map[string]bool)

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if set[item] {
				fmt.Println("cycle detected with: ", item)
			}
			if !seen[item] {
				seen[item] = true
				set[item] = true
				visitAll(m[item])
				order = append(order, item)
				delete(set, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	fmt.Println(keys)
	visitAll(keys)

	return order
}
