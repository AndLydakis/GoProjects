package main

import (
	"fmt"
	"hash/crc32"
	"os"
	"io"
)

func getHash(filename string) (uint32, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	h := crc32.NewIEEE()
	_, err = io.Copy(h, f)
	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil

}
func main() {
	h := crc32.NewIEEE()
	h.Write([]byte("Helloooooo"))
	v := h.Sum32()
	fmt.Println(v)

	h1, err := getHash("OReilly/test1.txt")
	if err != nil {
		return
	}
	h2, err := getHash("OReilly/test2.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)

}
