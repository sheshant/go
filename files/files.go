package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	defer f.Close()
	if err != nil {
		log.Default()
		panic(err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		log.Default()
		panic(err)
	}
	fmt.Println("name", fileInfo.Name())
	fmt.Println("size", fileInfo.Size())
	fmt.Println("mode", fileInfo.Mode())
	fmt.Println("modTime", fileInfo.ModTime())
	fmt.Println("isDir", fileInfo.IsDir())
	fmt.Println("sys", fileInfo.Sys())
	fmt.Println("mode", fileInfo.Mode())
	fmt.Println("mode", fileInfo.Mode().Perm())
	fmt.Println("mode", fileInfo.Mode().IsRegular())
	fmt.Println("mode", fileInfo.Mode().IsDir())
}
