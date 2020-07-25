// trunc.go


package main

import "fmt"

func main(){

    var fp_num float64

    fmt.Printf("Please enter a floating point number: ")
    num, err := fmt.Scan(&fp_num)

    trunc := int(fp_num)

    if (num == 1 && err == nil){
        fmt.Printf("Your integer is: %d\n", trunc)
    } else{
        fmt.Printf("Input Error\n")
    }
}
