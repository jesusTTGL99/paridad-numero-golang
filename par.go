package main

import "fmt"

func paridad(num int) bool{
	return num & 1 == 0;
}

func main() {
	var num int
	fmt.Printf("ingrese un numero natural: ")
	fmt.Scanln(&num)
	if paridad(num) {
		fmt.Println(num," es par")
	}else{
		fmt.Println(num," es impar")
	}
}
