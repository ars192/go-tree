package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(output io.Writer, path string, printFiles bool) error {
	iteration := 0
	dirTreeInternal(output, path, printFiles, iteration)
	return nil
}

func dirTreeInternal(output io.Writer, path string, printFiles bool, iteration int) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	lastElement := len(files)
	for i, file := range files {
		if i != lastElement-1 {
			fmt.Fprint(output, strings.Repeat("    ", iteration))
			fmt.Fprintln(output, "├───"+file.Name())
			if file.IsDir() == true {
				err = dirTreeInternal(output, path+"/"+file.Name(), printFiles, iteration+1)
				if err != nil {
					return err
				}
			}
		} else {
			fmt.Fprint(output, strings.Repeat("    ", iteration))
			fmt.Fprintln(output, "└───"+file.Name())
			if file.IsDir() == true {
				err = dirTreeInternal(output, path+"/"+file.Name(), printFiles, iteration+1)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}