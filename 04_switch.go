package main

import (
	"fmt"
	"time"
)

func main() {
	t:= time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	checkType := func(i interface{} ){
		switch t:= i.(type){
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Println("unknown %T", t)
		}
	}

	checkType("hey")
}