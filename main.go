package main

import "fmt"

func main() {
	var slice []string
	var slice2 []string
	slice = append(slice, "a")
	slice = append(slice, "b")
	slice = append(slice, "c")
	slice = append(slice, "d")
	slice2 = append(slice2, "A")
	slice2 = append(slice2, "B")
	slice2 = append(slice2, "C")
	slice2 = append(slice2, "D")
	var count int = len(slice)
	fmt.Println(count)
	count = 0
	for _, v := range slice {
		fmt.Println(v, slice2[count])
		count++
	}
}
