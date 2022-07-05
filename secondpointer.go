package main

import "fmt"

func main() {
	name := "original"
	setNameBroken(name, "qvault")
	fmt.Println(name)
}

func setNameBroken(ptr string, newName string) {
	ptr = newName
}

// output:
// original

// Pointers can be useful, but they can be just as dangerous as they are useful.
//  For example, if we use a pointer that has no value, the program will crash.
// That's why we always check to see if the error value is Nil before trying to print.

// If you print this pointer, you will see a memory address
// refer to thirdpointer
