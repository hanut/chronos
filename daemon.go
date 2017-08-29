package main

import (
	"fmt"
	"github.com/json-iterator/go"
	// "time"
)

func main() {
	fmt.Printf("Json Parse Test\n")
	user := Utd{}
	jsonData := []byte(`{"phone":"9717899743","balance":400.00,"cpm":4,"startedAt":"2017-08-29T23:16:15.917Z"}`)
	var tmp = jsoniter.Unmarshal(jsonData, &user)
	fmt.Println("Json Data-\n" + string(jsonData))
	fmt.Println("Parsed Data-\n" + user.String())
	fmt.Println("Balance : ", user.Balance)
	fmt.Println(tmp)
}
