package main

import (
	"fmt"
	"gopackages/citydigit"

	"github.com/huandu/xstrings"
)

func main() {
	fmt.Println(xstrings.Shuffle(citydigit.City()), citydigit.Digit())
}
