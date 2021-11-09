package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

func main() {

	gte := hello("")

	fmt.Println(aurora.Yellow(gte))

	gt := hello("John")
	fmt.Println(aurora.Yellow(gt))
}
