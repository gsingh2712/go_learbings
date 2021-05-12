package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	paths := os.Args[1:]

	if len(paths) == 0 {
		fmt.Println("Provide Directory Path")
		return
	}

	var dirs []byte
	for _, dir := range paths {
		files, err := ioutil.ReadDir(dir) // Reading contents inside the dir
		if err != nil {
			fmt.Println("Provide Correct Directories")
			return
		}

		dirs = append(dirs, dir...)
		dirs = append(dirs, '\n')

		for _, file := range files {

			if file.IsDir() { // Is content Directory
				dirs = append(dirs, '\t')
				dirs = append(dirs, file.Name()...) // File.Name() is a slice and this is how append slice to a slice
				dirs = append(dirs, '/', '\n')      // This is how to you append individual content to a slice
			}

		}
		dirs = append(dirs, '\n')
	}

	err := ioutil.WriteFile("dirstructure.txt", dirs, 0644)

	if err != nil {
		fmt.Println(err)
	}
}
