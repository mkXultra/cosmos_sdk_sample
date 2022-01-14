package main

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

	
type Person struct {
   firstName string 
   age int
}

type ori []byte
type Persons []Person
func main() {
	println("hello")
	addr := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	println(sdk.AccAddress(addr).String())
	println(sdk.AccAddress(addr).Bytes())
	println(sdk.AccAddress(addr))
	xx := ori(addr)
	jshon := []Person{{"jshon",1}}
	persons := Persons(jshon)
	println("-------")
	println(xx)
	println(persons)
	addr256 := make([]byte, 256)
	println(addr256)
	println(sdk.AccAddress(addr256))
}
