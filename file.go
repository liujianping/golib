package golib

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Directory(path string) (string, error) {
	fmt.Println("Call Direcotry:", path)
	if len(path) != 0 {
		return filepath.Abs(path)
	}
	return filepath.Abs(".")
}

func WalkDirectory(path string) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("walk fn failed: ", err)
			return err
		}
		fmt.Printf("walked file or dir: %q\n", path)
		return nil
	})
}
