package main

import (
	"fmt"
	"maps"
)

//! maps -> hash, objects, dict
func main() {

//? creating map
// m := make(map[string]string)
// ?setting an element
// m["name"] = "golang"
// m["tech"] = "backend"
//! if key does't exists in the map then it returns zero
// fmt.Println(m["name"], m["tech"], m["raj"])


// m := make(map[string]int);
// m["age"] = 30

// fmt.Println(m["age"], m["tech"])


// m := make(map[string]int)

// m["age"] = 20
// m["price"] = 30

// fmt.Println(len(m))

//? to delete
// delete(m, "price")
// fmt.Println(m)

//? to empty map totally
// clear(m)
// fmt.Println(m)



//! we can also declare map like this
// m := map[string]int{"price": 20, "age": 18, "watch": 300}
// fmt.Println(m)



//? Imp. Concept
// m := map[string]int{"age": 19, "price": 20}

//! here first returs us value and second returns us bool
// k,ok := m["price"]
// fmt.Println(k)

// if ok {
// 	fmt.Println("all ok")
// } else {
// 	fmt.Println("not ok")
// }



//? maps package
m1 := map[string]int{"age": 18, "price": 60}
m2 := map[string]int{"age": 18, "price": 30}
fmt.Println(maps.Equal(m1, m2))


}