package main

// What is a pointer?
// A pointer is a variable that stores the address of a value instead of the value itself.
// If you think of computer memory (RAM) as a JSON object, a pointer is like a key, and a regular variable is its value.

import "fmt"

func main() {
        // create a normal string variable
	name := "original"
        // pass in a pointer to the string variable using '&amp;amp;amp;'
	setName(&amp;amp;amp;name, "qvault")
	fmt.Println(name)
}

func setName(ptr *string, newName string) {
        // dereference the pointer so we can modify the value
        // and set the value to "qvault"
	*ptr = newName
}

// output :
// qvault

// As you can see, since we have a pointer to the address of the variable, we can change its value, even within the scope of another function.
// This wouldn't work if it wasn't a pointer value. Refer to the second pointer file.