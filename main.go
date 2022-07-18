package main

import (
	"encoding/json"
)

type Address struct {
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {

}