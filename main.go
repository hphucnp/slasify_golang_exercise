package main

import (
	"errors"
	"fmt"
	"os"
)

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

func printHelp() {
	fmt.Println("usage: \n\tgo run . <directory>")
}

func countCommentLines(dir string) error {
	// TODO: start your work here
	return errors.New(fmt.Sprintf(`
error:		not implemented. 
directory:	%s`, dir))
}
