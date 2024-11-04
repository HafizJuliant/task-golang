// main.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"task-golang/utils"
)

func main() {
    var n int
    if len(os.Args) > 1 {
        input, err := strconv.Atoi(os.Args[len(os.Args)-1])
        if err == nil {
            n = input
        } else {
            fmt.Println("Invalid argument. Using default value of 5 for n.")
            n = 5 // 
        }
    } else {
        fmt.Println("No argument provided. Using default value of 5 for n.")
        n = 5 // 
    }

    // Print results of utils functions
    fmt.Println("MagicSum:", utils.MagicSum(n))
    fmt.Println("MagicPow:", utils.MagicPow(n))
    fmt.Println("MagicOdd:", utils.MagicOdd(n))
    fmt.Println("MagicGrade:", utils.MagicGrade(n))
    fmt.Println("MagicName:", utils.MagicName(n))
    fmt.Println("MagicTria:", utils.MagicTria(n))

    
    fmt.Printf("Before MagicChange: %d\n", n)
    utils.MagicChange(&n)
    fmt.Printf("After MagicChange: %d\n", n)

   
    magicNum := utils.MagicNumber{Number: n}
    fmt.Printf("MagicNumber before Multiply: %+v\n", magicNum)
    magicNum.Multiply(2) // Example multiplying by 2
    fmt.Printf("MagicNumber after Multiply: %+v\n", magicNum)
}
