package main

import (
	"fmt"
	"time"
)

func main() {
	var s []string

	s = make([]string, 3)

	fmt.Printf("value: %v, length: %d, capacity: %d\n", s, len(s), cap(s))

	// assign
	s[1] = "assigned1"
	fmt.Printf("After assignment: %v\n", s)

	// append
	s = append(s, "assigned2", "assigned3")
	fmt.Printf("After append: %v\n", s)

	s2 := []string{"assigned4", "assigned5"}
	s = append(s, s2...)
	fmt.Printf("After append slice: %v\n", s)

	// copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Printf("After copy: %v\n", c)

	// slice c = [assigned1 assigned2 assigned3]
	slice1 := c[0:2] // include 0 to 1
	fmt.Printf("After slice1: %v\n", slice1)

	slice2 := c[2:] // include 2 to end
	fmt.Printf("After slice2: %v\n", slice2)

	slice3 := c[:2] // exclude 2
	fmt.Printf("After slice3: %v\n", slice3)

	// map
	m := make(map[string]string)
	m["name"] = "John"
	m["age"] = "30"
	fmt.Printf("After map: %v\n", m)

	fmt.Printf("Get value of name key: %v\n", m["name"])

	// delete(m, "name")
	// fmt.Printf("After delete: %v\n", m)

	// clear(m)
	// fmt.Printf("After clear: %v\n", m)

	// struct
	type Person struct {
		name string
		age  int
	}

	p := Person{name: "John", age: 30}
	fmt.Printf("After struct: %v\n", p)

	// test preallocate time
	num := 1000000
	testSlice := []int{}
	testSlice2 := make([]int, num)

	fmt.Printf("Total time without preallocate: %v\n", testPreallocate(testSlice, num))
	fmt.Printf("Total time with preallocate: %v\n", testPreallocate(testSlice2, num))

	// pointer
	var x int = 10
	var y int = 20

	var xPtr *int = &x
	var yPtr *int = &y

	fmt.Printf("xPtr: %v, yPtr: %v\n", xPtr, yPtr)
}

func testPreallocate(testSlice []int, num int) time.Duration {
	start := time.Now()

	for len(testSlice) < num {
		testSlice = append(testSlice, 1)
	}

	return time.Since(start)
}
