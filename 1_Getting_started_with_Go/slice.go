//slice.go

package main

import (
    "fmt"
    "sort"
    "strconv"
)

func main() {
    var arr []int = make([]int, 3)
    var input string
    fmt.Println("Please enter an interger(X to exit):")
    for true {
        fmt.Scanln(&input) 
        if input == "X"{
            break
        }
        num,err:=strconv.Atoi(input)
        if err != nil {
            fmt.Println("Wrong input")
            continue
        }
        arr = append(arr, num)
        sort.Ints(arr[:])
        fmt.Println(arr)
        fmt.Println("Please again enter an interger(X to exit):")
	}
}
