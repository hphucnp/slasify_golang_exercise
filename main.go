package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// FileStats represents the total lines, inline comments, and block comments in a file
type FileStats struct {
	TotalLines     int
	InlineComments int
	BlockComments  int
}

// Entry point of the program
func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		printHelp()
	} else {
		dir := args[0]
		if err := countCommentLines(dir); err != nil {
			fmt.Println(err)
		}
	}
}

// Prints the usage guide if the user inputs incorrectly
func printHelp() {
	fmt.Println("usage: \n\tgo run . <directory>")
}

// countCommentLines is the core logic to process each C/C++ file in the directory
func countCommentLines(dir string) error {
	// Verify that the directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("error: directory does not exist: %s", dir))
	}

	// Walk the directory to collect all C/C++ source files
	files, err := collectSourceFiles(dir)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errors.New(fmt.Sprintf("error: no C/C++ source files found in directory: %s", dir))
	}

	// Process and print the comment line statistics for each file
	stats, err := processFiles(files)
	if err != nil {
		return err
	}

	// Print the results in the required format
	printResults(stats)

	return nil
}

// collectSourceFiles walks through the directory and gathers all C/C++ files
func collectSourceFiles(dir string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isCSourceFile(info.Name()) {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking through directory: %s", err)
	}

	sort.Strings(files)
	return files, nil
}

// isCSourceFile determines if the given file is a C or C++ source/header file
func isCSourceFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".c" || ext == ".cpp" || ext == ".h" || ext == ".hpp"
}

// processFiles calculates the comment line statistics for each file
func processFiles(files []string) (map[string]FileStats, error) {
	stats := make(map[string]FileStats)

	for _, file := range files {
		fileStats, err := analyzeFile(file)
		if err != nil {
			return nil, fmt.Errorf("error processing file %s: %v", file, err)
		}
		stats[file] = fileStats
	}

	return stats, nil
}

// analyzeFile reads a file and counts its total, inline, and block comment lines
func analyzeFile(file string) (FileStats, error) {
	f, err := os.Open(file)
	if err != nil {
		return FileStats{}, err
	}
	defer f.Close()

	stat := FileStats{}
	inBlockComment := false
	blockStart := regexp.MustCompile(`/\*`)
	blockEnd := regexp.MustCompile(`\*/`)
	inlineComment := regexp.MustCompile(`//`)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		stat.TotalLines++

		if inBlockComment {
			stat.BlockComments++
			if blockEnd.MatchString(line) {
				inBlockComment = false
			}
			continue
		}

		if blockStart.MatchString(line) {
			stat.BlockComments++
			if !blockEnd.MatchString(line) {
				inBlockComment = true
			}
		} else if inlineComment.MatchString(line) {
			stat.InlineComments++
		}
	}

	if err := scanner.Err(); err != nil {
		return FileStats{}, err
	}

	return stat, nil
}

// printResults outputs the file statistics in the required format
func printResults(stats map[string]FileStats) {
	// Sort the file paths to ensure alphabetical output
	files := make([]string, 0, len(stats))
	for file := range stats {
		files = append(files, file)
	}
	sort.Strings(files)

	// Print the results
	for _, file := range files {
		stat := stats[file]
		// Print formatted output with aligned columns
		fmt.Printf("%-40s total: %4d    inline: %3d    block: %3d\n", file, stat.TotalLines, stat.InlineComments, stat.BlockComments)
	}
}
