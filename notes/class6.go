/*Review Code 
3 Challenges
- assign variables correctly
- different ways to print Sprint, Fprint, println
- express myself in exchanges values between variables
- setting up golint (45 minutes)
- how read Go test for Exercism problems
- practice creating arrays then slices of arrays
*/

//Cheatsheat

// Variables
package main

var fiveIntegers [5]int
fiveIntegers[4] = 100

// For loo
for i := 7;  i <=9 i++ {

}
// Slices
mySlice := make([][]string, 3)
var mySliceCopy
copy(mySlice, mySliceCopy)

// Maps
m := make(map[string]int) // add get delete operations
m["values1"] = 7
m["values2] = 8

v1 := m["values"]
delete(m, "values1")

_, exists := m["values"]

otherMap := map[string]int{"id": 1, "count":12}

// Struct

type person struct {
	id int
	age int
	name string
}

func NewPerson(personName string, personId int, personAge int) *person {
	newPerson:= person{name: personName, age: personAge, id:personId}
	// newPerson.age = 2
	return $newPerson
}

// Ask your buddy where your friend is == pointer

// & gives you the address of a variable
// * dereference the pointer
	// 