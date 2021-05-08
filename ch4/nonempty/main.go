package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for idx, s := range strings {
		fmt.Printf("%v,%v\n", idx, s)
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}
