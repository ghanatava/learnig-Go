package main

import "fmt"

func sum(nums ...int){
     fmt.Println("nums= ",nums)
     total := 0

     for _,i := range nums{
	total+= i
     }
     fmt.Println("sum = ",total)
}

func main(){
     sum(1,2)
     sum(2,34,5)
     nums := []int{2,3,4,5}
     sum(nums...)  //the following ... passes elements sperately #unpacking
}
