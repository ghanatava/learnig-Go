package main

import "fmt"

func vals()(int,int){
     return 3,4
}

func main(){
     a,b := vals()
     fmt.Printf("a= %v and b= %v\n",a,b)

     c,_ := vals()
     fmt.Printf("c= %v\n",c)
}
