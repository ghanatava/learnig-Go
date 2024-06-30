package main

import (
	"fmt"
	"log"

	"github.com/ghanatava/learning-Go/protocol_buffers/internal"
	"google.golang.org/protobuf/proto"
)

func main() {
	p := &internal.Person{
		Name: "elliot",
		Age:  24,
		Socialfollowers: &internal.SocialFollowers{
			Twitter: 3200,
			Youtube: 2500,
		},
	}

	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal("Marshaling error", err)
	}
	fmt.Println(data)

	p1 := &internal.Person{}
	err = proto.Unmarshal(data, p1)
	if err != nil {
		log.Fatal("Unmarshaling error", err)
	}
	fmt.Println(p1)
}
