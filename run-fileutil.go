package main

// FileCopy utility

import (
	"fileutil"
	"fmt"
)

func main() {
	err := fileutil.FileCopy("my-test-file.txt", "my-test-file-copy.txt")
	if err != nil {
		fmt.Println(err)
	}
}
