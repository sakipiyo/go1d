package main

import "fmt"

func main() {
	intarray := [6]int{98, 125, 232, 147, 486, 500}
	slice13 := intarray[1:4]
	slice14 := intarray[1:5]
	slice0 := intarray[:2]
	slice4 := intarray[3:]
	slice13[2] -= 100

	fmt.Printf("slice13は%d\n", slice13)
	fmt.Printf("slice14は%d\n", slice14[1:5])
	fmt.Printf("slice13[1]は%d\n", slice13[1])
	fmt.Printf("slice13[0]は%d\n", slice13[0])
	fmt.Printf("slice0の要素数は%d\n", len(slice0))
	fmt.Printf("slice0[1]は%d\n", slice0[1])
	fmt.Printf("slice4[0]の値は%d\n", slice4[0])
	fmt.Printf("今やiniarray[3]の値は%d\n", intarray[3])
}
