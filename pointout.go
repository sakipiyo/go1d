package main

import "fmt"

func main() {

	inta := 5
	adinta := &inta

	*adinta = 7

	fmt.Println("*adintaを変更：")
	fmt.Printf("inta の値は %d\n", inta)
}
