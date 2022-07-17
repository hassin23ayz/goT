package main

import "fmt"

func main(){
	
	s:= make([]string, 3)
	fmt.Println("emp:", s)
	fmt.Println("len", len(s))

	s = append(s, "4")
	s = append(s, "5")
	s = append(s, "6")

	fmt.Println("emp:", s)	

	s[0] = "a"
	s[1] = "b"

	fmt.Println("emp:", s)

	
}