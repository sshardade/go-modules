package main

import (
	"fmt"
	"github.com/sshardade/go-modules/greet"
	"rsc.io/quote/v3"
)

func main() {
	fmt.Println("First code using go modules...")
	greet.GreetQuotes("Suraj Hardade")
	fmt.Println(quote.HelloV3())
}
