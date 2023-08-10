package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, _ := rpc.Dial("tcp", "localhost:8081")

	if err := client.Call("Handler.GetUsers", 1, nil); err != nil {
		fmt.Printf("Error:1 user.GetUsers() %+v", err)
	} else {
		fmt.Printf("user found")
	}
}
