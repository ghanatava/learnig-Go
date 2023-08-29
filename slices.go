package main

import (
    "fmt"
    "strconv"
)

func main(){
    var s []string
    fmt.Println("uninit ",s,s==nil,len(s)==0)

    s=make([]string,3)
    fmt.Println("emp ",s,"len ",len(s),"cap ",cap(s))

    for i := 0; i<3;i++{
	s[i]=strconv.Itoa(i)
    }

    fmt.Println(s)
   
    s=append(s,"a")
    s=append(s,"b")
    
    fmt.Println(s)

    l := s[2:5]
    fmt.Println("l= ",l) 
    twoD := make([][]int,3)
    for i := 0; i<3;i++{
	inner_len := i+1
        twoD[i] = make([]int,inner_len)
        
	for j := 0; j<inner_len; j++{
	    twoD[i][j] = i+j
	}
    }
    fmt.Println("Twod ",twoD)
}
