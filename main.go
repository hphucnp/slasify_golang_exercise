package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type FileStats struct {
	TotalLines     int
	InlineComments int
	BlockComments  int
}

type LanguageSpec struct {
	inlineCommentStart string
	blockCommentStart       string
	blockCommentEnd         string
	lineContinuation        string
	escapeCharacter         string
	specialCharsToEscape    []string
	stringDelimiter         []string
	fileExtensions          []string
}

var languageSpecs = map[string]LanguageSpec{
    "C/C++": {
        inlineCommentStart: "//",
        blockCommentStart:  "/*",
        blockCommentEnd:    "*/",
		lineContinuation:   "\\",
		escapeCharacter:    "\\",
		specialCharsToEscape: []string{"\\", "\"", "'", "a", "b", "f", "n", "r", "t", "v"},
        stringDelimiter:    []string{"\"", "'"},
        fileExtensions:     []string{".c", ".cpp", ".h", ".hpp"},
    },
    // Add more language specifications as needed
}

var activeLanguage LanguageSpec

// Entry point of the program
func main() {
	configured_lang := os.Getenv("ACTIVE_LANGUAGE")
	if configured_lang == "" {
		configured_lang = "C/C++"
	}
	var ok bool
	activeLanguage, ok  = languageSpecs[configured_lang]
	if !ok {
		fmt.Println("error: unsupported language: ", configured_lang)
		return
	}
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
	fmt.Print()
}

// countCommentLines is the core logic to process each C/C++ file in the directory
func countCommentLines(dir string) error {
	// Verify that the directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("error: directory does not exist: %s", dir)
	}

	// Walk the directory to collect all C/C++ source files
	files, err := collectSourceFiles(dir)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return fmt.Errorf("error: no C/C++ source files found in directory: %s", dir)
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
	for _, cExt := range activeLanguage.fileExtensions {
		if ext == cExt {
			return true
		}
	}
	return false
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
	isInBlockComment := false
	isInInlineComment := false
	isInString := false
	currentStringDelimiter := ""
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		stat.TotalLines++

		isCountedAsBlockCommment := false
		lineContinued := false
		if (len(line) > 0) {
			lineContinued = line[len(line)-1] == activeLanguage.lineContinuation[0]
		}
			
		i := 0
		if line == "" {
			if isInInlineComment {
				stat.InlineComments++
				isInInlineComment = false	
			}
			if isInBlockComment {
				stat.BlockComments++
			}
			continue
		}
		for i < len(line) {
			if isInString {
				// Check if the string ends on this line
				if strings.HasPrefix(line[i:], currentStringDelimiter) {
					isInString = false
					i += len(currentStringDelimiter)
					currentStringDelimiter = ""
					continue
				}

				// Check if the character is an escape character
				if strings.HasPrefix(line[i:], activeLanguage.escapeCharacter) {
					i += 2
					continue
				}
			}

			if isInBlockComment {
				if !isCountedAsBlockCommment {
					stat.BlockComments++
					isCountedAsBlockCommment = true
				}
				// Check if the block comment ends on this line
				if strings.HasPrefix(line[i:], activeLanguage.blockCommentEnd) {
					isInBlockComment = false
					i += len(activeLanguage.blockCommentEnd)
					continue
				}
			}

			if isInInlineComment {
				// The rest of the line is a comment
				stat.InlineComments++
				if !lineContinued {
					isInInlineComment = false
				}
				break
			}

			// Check if the line is a string
			start_string := false
			for _, delimiter := range activeLanguage.stringDelimiter {
				if strings.HasPrefix(line[i:], delimiter) && !isInBlockComment && !isInInlineComment && !isInString {
					isInString = true
					start_string = true
					currentStringDelimiter = delimiter
					i += len(delimiter)
					break
				}
			}
			if start_string {
				continue
			}
			// Check if the line is a block comment
			if strings.HasPrefix(line[i:], activeLanguage.blockCommentStart) && !isInString && !isInBlockComment && !isInInlineComment {
				isInBlockComment = true
				i += len(activeLanguage.blockCommentStart)
				if !isCountedAsBlockCommment {
					stat.BlockComments++
					isCountedAsBlockCommment = true
				}
				continue
			}

			// Check if the line is an inline comment
			if strings.HasPrefix(line[i:], activeLanguage.inlineCommentStart) && !isInString && !isInBlockComment {
				stat.InlineComments++
				isInInlineComment = true
				if !lineContinued {
					isInInlineComment = false
				}
				break
			}

			// Move to the next character
			i++
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
