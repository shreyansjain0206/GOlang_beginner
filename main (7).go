/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/
package main

import "fmt"

// Function with parameters and return value
func add(x, y int) int {
    return x + y
}

func main() {
    result := add(3, 5)
    fmt.Println("Sum:", result)
}