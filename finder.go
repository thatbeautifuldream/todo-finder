package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type Comment struct {
	File    string
	Line    int
	Comment string
}

func findComments(dir string) []Comment {
	var results []Comment

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".go") {
			fileResults := scanFile(path)
			results = append(results, fileResults...)
		}

		return nil
	})

	return results
}

func scanFile(path string) []Comment {
	var results []Comment

	file, err := os.Open(path)
	if err != nil {
		return results
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		if strings.Contains(line, "TODO") || strings.Contains(line, "FIXME") {
			results = append(results, Comment{
				File:    path,
				Line:    lineNumber,
				Comment: strings.TrimSpace(line),
			})
		}
	}

	return results
}
