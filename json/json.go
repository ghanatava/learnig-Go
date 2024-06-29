package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID          int      `json:"user_id"` //rename ID to user_id
	Name        string   `json:"name,omitempty"`
	Age         int      `json:"age"`
	Password    string   `json:"-"`     //ignore password
	Permissions []string `json:"roles"` //rename permission to roles
}

func main() {
	u := User{
		ID:          1,
		Name:        "Ashura",
		Age:         23,
		Password:    "mysecret",
		Permissions: []string{"admin", "group-member"},
	}

	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(string(b))

	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID %v, Name: %v, Age: %v \n", u.ID, u.Name, u.Age)

	f, err := os.Create("output.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	err = encoder.Encode(u)
	if err != nil {
		panic(err)
	}

	var u2 User
	f2, err := os.Open("output.json")
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	decoder := json.NewDecoder(f2)
	err = decoder.Decode(&u2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID %v, Name: %v, Age: %v \n", u2.ID, u2.Name, u2.Age)

}
