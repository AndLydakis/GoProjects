package main

import "fmt"
import (
	"os"
	"path/filepath"
)

func main() {
	file, err := os.Open("OReilly/test.txt")
	if err != nil {
		fmt.Println("Error 1")
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error 2")
		return
	}
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("Error 3")
		return
	}
	str := string(bs)
	fmt.Println(str)

	dir, err := os.Open(".")
	if err != nil{
		fmt.Println("Error 4")
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi :=range fileInfos{
		fmt.Println(fi)
	}

	filepath.Walk(".", func(path string, info os.FileInfo, err error)error{
		fmt.Println(path)
		return nil
	})
}
