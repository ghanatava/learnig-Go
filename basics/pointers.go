package main

import "fmt"

func zeroval(val int){
	val = 0
	fmt.Println("i = ",val)
}

func zeroptr(ptr *int){
	*ptr = 0
	fmt.Println("i = ",*ptr)
}

func main(){
	i := 1
	fmt.Println("i = ",i)
	
	zeroval(i)
	fmt.Println("i = ",i)
	
	zeroptr(&i)
	fmt.Println("i = ",i)
}
