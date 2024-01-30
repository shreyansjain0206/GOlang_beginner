/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/
package main

import "fmt"

func main() {
    day := "Monday"

    switch day {
    case "Monday":
        fmt.Println("It's Monday.")
    case "Tuesday":
        fmt.Println("It's Tuesday.")
    default:
        fmt.Println("It's another day.")
    }
}