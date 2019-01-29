package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dineshbalachandran/github"
	"github.com/dineshbalachandran/shapes"
)

func main() {
	intset()
}

func intset() {
	var s, t IntSet

	s.Add(100)
	s.Add(201)

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
