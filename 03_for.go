package main

import "fmt"

func main(){

	for i :=1 ; i<=7; i++ {
		fmt.Println(i)
		break
	}

	for n:=0; n<=5; n++{
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}