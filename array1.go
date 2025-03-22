package main

import (
	"fmt"
)

func main() {
	var array1 [5]int
	array1[0] = 10
	array1[1] = 20
	array1[2] = 30
	array1[3] = 40
	array1[4] = 50
	fmt.Println(array1)

	array2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(array2)

	array3 := [...]int{10, 20, 30, 40, 50}
	fmt.Println(array3)

	array4 := [5]int{1: 10, 3: 30}
	fmt.Println(array4)

	array5 := [5]int{0: 10, 4: 50}
	fmt.Println(array5)

	//Trabajemos con slices
	slice1 := []int{10, 20, 30, 40, 50}
	fmt.Println(slice1)

	slice2 := array1[1:4]
	fmt.Println(slice2)

	slice3 := array1[:4]
	fmt.Println(slice3)

	slice4 := array1[1:]
	fmt.Println(slice4)

	//Trabajemos con map
	map1 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(map1)

	map2 := map[string]int{
		"one":   1, // comma is necessary
		"two":   2, // comma is necessary
		"three": 3, // comma is necessary
	}
	fmt.Println(map2)

	map3 := make(map[string]int)
	map3["uno"] = 1  // comma is necessary
	map3["dos"] = 2  // comma is necessary
	map3["tres"] = 3 // comma is necessary

	fmt.Println(map3)

}
