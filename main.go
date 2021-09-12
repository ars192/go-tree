package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

const MiddleEdge = "├───"
const Edge = "│\t"
const EndEdge = "└───"
const Tab = "\t"
const Empty = "(empty)"

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
	var a []string
	err := dirTreeInternal(output, path, printFiles, a)
	if err != nil {
		return err
	}
	return nil
}

func dirTreeInternal(output io.Writer, path string, printFiles bool, tabs []string) error {
	files, err := ioutil.ReadDir(path)
	var newFiles []os.FileInfo
	if err != nil {
		return err
	}
	if printFiles == false {
		for _, v := range files {
			if v.IsDir() {
				newFiles = append(newFiles, v)
			}
		}
	} else {
		newFiles = files
	}

	lastElement := len(newFiles)
	for i, file := range newFiles {
		if i != lastElement-1 {
			for _, v := range tabs {
				_, err := fmt.Fprint(output, v)
				if err != nil {
					return err
				}
			}
			var b []string
			b = append(b, tabs...)

			if file.IsDir() {
				_, err := fmt.Fprintln(output, MiddleEdge+file.Name())
				if err != nil {
					return err
				}
				b = append(b, Edge)
				err = dirTreeInternal(output, path+"/"+file.Name(), printFiles, b)
				if err != nil {
					return err
				}
			} else {
				if printFiles {
					size := strconv.Itoa(int(file.Size()))
					if size == "0" {
						size = Empty
					} else {
						size = "(" + size + "b)"
					}
					_, err := fmt.Fprintln(output, MiddleEdge+file.Name()+" "+size)
					if err != nil {
						return err
					}
				} else {
					continue
				}
			}
		} else {
			for _, v := range tabs {
				_, err := fmt.Fprint(output, v)
				if err != nil {
					return err
				}
			}
			var b []string
			b = append(b, tabs...)
			if file.IsDir() {
				_, err := fmt.Fprintln(output, EndEdge+file.Name())
				if err != nil {
					return err
				}
				b = append(b, Tab)
				err = dirTreeInternal(output, path+"/"+file.Name(), printFiles, b)
				if err != nil {
					return err
				}
			} else {
				if printFiles {
					size := strconv.Itoa(int(file.Size()))
					if size == "0" {
						size = "(empty)"
					} else {
						size = "(" + size + "b)"
					}
					_, err := fmt.Fprintln(output, EndEdge+file.Name()+" "+size)
					if err != nil {
						return err
					}
				} else {
					continue
				}
			}
		}
	}
	return nil
}
