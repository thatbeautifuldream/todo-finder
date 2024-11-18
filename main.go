package main

import (
	"flag"
	"fmt"
	"os"
)


func main() {
	dir := flag.String("dir", ".", "Directory to scan")
	flag.Parse()

	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exist\n", *dir)
		os.Exit(1)
	}

	results := findComments(*dir)

	if len(results) == 0 {
		fmt.Println("No TODO or FIXME comments found!")
	} else {
		fmt.Println("Found comments:")
		for _, res := range results {
			fmt.Printf("%s:%d: %s\n", res.File, res.Line, res.Comment)
		}
	}
}
