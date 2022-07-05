package main

import "fmt"

func main() {
	newString := ""
	newStringPointer := &amp;amp;amp;newString,
	fmt.Println(newStringPointer)
}

// output in wrong expect:
// 0xc00000e1e0

// When should we use pointers?
// Most of the time you use pointers, it must be because of one of the following:
// 1. A function that changes one of its parameters.
// When we call a function that takes a pointer as a parameter, we expect our variable to be changed.
//  If you're not changing the variable in your function, then you probably shouldn't be using pointers.
// 2. Better performance
// If you have a string that contains an entire novel in memory, 
// copying this variable every time it is passed to a new function is very expensive.
// . It may be worthwhile to pass a pointer instead, which saves CPU and memory.
//  However, doing so comes at the cost of readability, so only do this optimization if necessary.
// 3. You need Nil option.
// Sometimes a function needs to know what the value of something is, it also needs to know if it exists or not. 
// We usually use this when reading JSON to know if a field exists or not. For example, if it is a JSON object.