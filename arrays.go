package main

import "fmt"

func main(){
    var a [5]int
    fmt.Println("emp: ",a)

    a[4]=100
    fmt.Printf("a[4]=%v\n",a[4])
    
    b := [5]int {1,2,3,4,5}
    fmt.Println("b= ",b)

    var twod [2][5]int 
    for i := 0;i<2;i++{
	for j := 0; j<5; j++{
            twod[i][j] = i+j
	}
    }
    fmt.Println("Two dimensional: ",twod)
}

