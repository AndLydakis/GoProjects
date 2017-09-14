package main

import "fmt"
import "strings"

func main() {
	fmt.Println("Contains")
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println("Count")
	fmt.Println(strings.Count("test", "t"))
	fmt.Println("HasPrefix")
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println("Index")
	fmt.Println(strings.Index("test", "te"))
	fmt.Println("Join")
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat")
	fmt.Println(strings.Repeat("a", 5))
	fmt.Println("Replace")
	fmt.Println(strings.Replace("aaaaaaaaaa", "a", "b", -1))
	fmt.Println("Split")
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower")
	fmt.Println(strings.ToLower("AAAAAAAAAAAAAAAA"))
	fmt.Println("ToUpper")
	fmt.Println(strings.ToUpper("AAAAAAAAAAAAAAAA"))
}
