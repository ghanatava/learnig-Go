package main

import (
    "fmt"
    "time"
)

func main(){
    i := 2
    switch i{
    case 1:
	fmt.Println(i)	
    case 2:
	fmt.Println(i)
    case 3:
	fmt.Println(i)
    }
    
    switch time.Now().Weekday(){
    case time.Saturday, time.Sunday:
	fmt.Println("Weekend")
    default:
	fmt.Println("Weekday")
    }
    
    whatami := func(i interface{}){
	switch t := i.(type){
	    case bool:
		fmt.Println("bool")
	    case int:
		fmt.Println("int")
	    default:
		fmt.Printf("Don't know the type %T\n",t)
	}
    }
    whatami(true)
    whatami(1)
    whatami("hey")
}


