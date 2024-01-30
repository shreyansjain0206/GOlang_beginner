/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/

package main

import "fmt"

// Define a struct as a class
type Car struct {
    Brand  string
    Model  string
    Year   int
    Engine string
}

func main() {
    // Create an instance of the Car struct
    myCar := Car{
        Brand:  "Toyota",
        Model:  "Camry",
        Year:   2022,
        Engine: "V6",
    }

    // Access fields of the struct
    fmt.Printf("My car is a %d %s %s with a %s engine.\n", myCar.Year, myCar.Brand, myCar.Model, myCar.Engine)
}