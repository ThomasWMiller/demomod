// fileutil package provides utilities for working with files
package fileutil

// FileCopy utility

import (
	"fmt"
	"io"
	"os"
)

// fileutil.FileCopy(sourceFile, destFile)
// Copies sourceFile to destFile.
// Reports copy to the console.
// Defers and reports closing of files.
func FileCopy(sourceFile, destFile string) error {
	src, err := os.Open(sourceFile)
	if err != nil {
		return err
	}

	defer func() {
		fmt.Println("closing", sourceFile)
		src.Close()
	}()

	dst, err := os.Create(destFile)
	if err != nil {
		return err
	}

	defer func() {
		fmt.Println("closing", destFile)
		src.Close()
	}()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	fmt.Printf("%s copied to %s\n", sourceFile, destFile)
	return nil
}
