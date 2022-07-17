package main 

import "fmt"

func main() {
	
	var a[5]int
	fmt.Println("emp:", a)

	fmt.Println("len:", len(a))
	a[0] = 5

	fmt.Println("emp:", a)

}