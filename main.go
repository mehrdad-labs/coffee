package main

import (
	"fmt"
	"math/rand"
)

var ilovecoffee string
var single string
var double string
var latte string

func init() {
	ilovecoffee =`                (
                 )
I ❤️ Coffee -> c[]
`

	single = `
  .-=-.
 ,|` + "`" + `~'|
 ` + "`" + `|   |  Single Espresso
   ` + "`" + `~'
`
	double = `
  .-~~-.
,|` + "`" + `-__-'|
||      |
` + "`" + `|      |  Double Espresso
  ` + "`" + `-__-'
`

	latte = `
  .=%%=.
,|` + "`" + `=%%='|
||      |
` + "`" + `|      |  Double Latte with Foam
  ` + "`" + `-__-'
`

}

func main() {

	num := rand.Intn(256)
	dv := num % 3

	if dv == 0 {

		fmt.Print(single)
		fmt.Print(ilovecoffee)

	} else if dv == 1 {

		fmt.Print(double)
		fmt.Print(ilovecoffee)

	} else { //* == 2

		fmt.Print(latte)
		fmt.Print(ilovecoffee)

	}

}
