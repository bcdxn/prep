package main

import (
	"fmt"

	ds "github.com/bcdxn/leetcode/fundamentals/data_structures"
)

func main() {
	t := ds.NewTrie()

	t.Insert("apple")
	t.Insert("ape")

	fmt.Println("contains apple: ", t.Search("apple"))
	fmt.Println("contains ape: ", t.Search("ape"))
	fmt.Println("contains app: ", t.Search("app"))

	fmt.Println("contains word prefixed with ap", t.StartsWith("ap"))
	fmt.Println("contains word prefixed with el", t.StartsWith("el"))
}
