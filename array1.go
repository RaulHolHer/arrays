package main

import (
	"fmt"
	"sort"
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

	map4 := make(map[string]int)
	map4["R"] = 54
	map4["P"] = 53
	map4["M"] = 22
	fmt.Println(map4)

	// crear un slice para ordenar las clave de un map
	keys := make([]string, 0, len(map4))
	for k := range map4 {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	// ordenar el slice
	sort.Strings(keys) // uncomment this line to sort the keys
	fmt.Println(keys)

	// imprimir el map ordenado
	for _, k := range keys {
		fmt.Println(k, map4[k])
	}
	// bucles
	for i := 0; i < 5; i++ {
		//ftm.Println(i)
		fmt.Print(i, " ") // imprime en la misma linea
	}
	fmt.Println("") // imprime una linea en blanco

	// bucle de array1. len(array1) devuelve la longitud del array
	for indice := 0; indice < len(array1); indice++ {
		fmt.Println(array1[indice])
	}
	// bucle utilizando el range (rango)
	for indice, valor := range array1 {
		fmt.Println(indice, valor)
	}
	// bucle utilizando el range (rango) sin utilizar el indice
	for _, valor := range array4 {
		fmt.Println(valor)
	}

}
