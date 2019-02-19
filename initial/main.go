package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dineshbalachandran/github"
	"github.com/dineshbalachandran/shapes"
	"github.com/dineshbalachandran/tempconv"
)

type sortface struct {
	t []int
}

func (s sortface) Len() int           { return len(s.t) }
func (s sortface) Less(i, j int) bool { return s.t[i] < s.t[j] }
func (s sortface) Swap(i, j int)      { s.t[i], s.t[j] = s.t[j], s.t[i] }

func main() {

	s := sortface{t: []int{1, 1, 3, 1}}
	fmt.Printf("%t\n", isPalindrome(s))
}

func testflag() {
	var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")
	flag.Parse()
	fmt.Println(*temp)
}

func topsort() {
	// prereqs maps computer science courses to their prerequisites.
	var prereqs = map[string][]string{
		"algorithms":     {"data structures"},
		"calculus":       {"linear algebra"},
		"linear algebra": {"calculus"},
		"abc":            {"xyz"},
		"xyz":            {"abc"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}

	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func bar(s string) string {
	return "bar"
}

func intset() {
	var s, t IntSet

	s.AddAll(100, 201)

	t.Add(900)

	s.UnionWith(&t)

	fmt.Println(&s)
	fmt.Println(s.Len())

	s.Remove(201)

	fmt.Println(&s)
	fmt.Println(s.Len())

	x := s.Copy()
	fmt.Println(x)

	var y IntSet
	y.Add(100)
	y.Add(300)

	s.DifferenceWith(&y)
	fmt.Println(&s)

	s.SymmetricDifference(&y)
	fmt.Println(&s)

	fmt.Println(s.Elems())

	s.Clear()
	fmt.Println(&s)
	fmt.Println(s.Len())
}

func searchissues() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	now := time.Now()
	oneMonth := now.AddDate(0, -1, 0)
	oneYear := now.AddDate(-1, 0, 0)
	category := ""
	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(oneMonth):
			category = "<1 month"
		case item.CreatedAt.After(oneYear):
			category = "<1 year"
		default:
			category = ">1 year"
		}

		fmt.Printf("#%-5d %8s %9.9s %.55s\n",
			item.Number, category, item.User.Login, item.Title)
	}
}

func printargs() {
	for i, a := range os.Args {
		fmt.Println(i, a)
	}
}

func printarea() {
	shapes := []shapes.Shape{shapes.Circle{X: 0.0, Y: 0.0, R: 5.0},
		shapes.Rectangle{X1: 0, Y1: 0, X2: 2, Y2: 2}}

	for _, s := range shapes {
		fmt.Println(s.Area())
	}
}
