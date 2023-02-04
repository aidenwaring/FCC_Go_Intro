package constants

import "fmt"

const myConst int = 20

// Counter for enumerated constants
// const a = iota

const ( // Value of iota is scoped to this constant block
	x = iota
	y = iota
	z = iota
)

func Constants() {
	fmt.Printf("%v, %T\n", myConst, myConst)

	// const j :=  20  --- Error, see correct const infered type declaration below
	const j = 42

	fmt.Println(j)
	fmt.Println(x, y, z)
}