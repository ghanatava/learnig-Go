package main

import (
    "fmt"
    "maps"  //only for using maps.Equal you can create maps without importing
)

func main(){
    m := make(map[string]int)
    
    m["k1"]=7
    m["k2"]=13

    fmt.Println("map ",m)
    fmt.Println("v1= ",m["k1"])
    fmt.Println("v2= ",m["k2"])
    fmt.Println("len= ",len(m))
 
    clear(m)

    _,v2 := m["k2"]
    fmt.Println(v2)

    n := map[string]int {"foo":1,"bar":2}
    fmt.Println("map ",n)
    
    n2 := map[string]int {"foo":1,"bar":2}
    fmt.Println("map2 ",n2)

    if maps.Equal(n,n2) {
	fmt.Println("n == n2")
    } else{
        fmt.Println("n != n2")
    }
    
}
