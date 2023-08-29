package main

import "fmt" 

//note range iterates over data structures instead of creating a range 
//so it is more like a for each loop in java than range in python 

func main(){
    nums := []int{2,3,4}
    sum := 0
    for _,num := range nums{
	sum+=num
    }
    fmt.Println("sum = ",sum)
    
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

     kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
	

    for k := range kvs {
        fmt.Println("key:", k)
    }	
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}

