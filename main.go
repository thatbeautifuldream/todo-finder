package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "v1.0.0"

func main() {
	versionFlag := flag.Bool("version", false, "Print the version and exit")
	versionShortFlag := flag.Bool("v", false, "Print the version and exit")
	dir := flag.String("dir", ".", "Directory to scan")
	flag.Parse()

	if *versionFlag || *versionShortFlag {
		fmt.Println("todo-finder version:", version)
		os.Exit(0)
	}

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
