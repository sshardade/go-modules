package greet

import (
	"fmt"
	"time"

	"rsc.io/quote"
)

func GreetQuotes(name string) {
	tNow := time.Now()

	switch tHour := tNow.Hour(); {
	case tHour >= 5 && tHour < 12:
		fmt.Printf("Good Morning, %v!!!\n", name)
	case tHour >= 12 && tHour < 16:
		fmt.Println("Good Afternoon, ", name, "!!!'")
	case tHour >= 16 && tHour < 22:
		fmt.Printf("Good Evening, %v!!!\n", name)
	default:
		fmt.Printf("Good Night, %v!!!\n", name)
	}

	fmt.Println(quote.Go())

}
