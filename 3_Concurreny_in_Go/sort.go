/******** 
Write a program to sort an array of integers. The program should partition the array 
into 4 parts, each of which is sorted by a different goroutine. Each partition 
should be of approximately equal size. Then the main goroutine should merge the 4 
sorted subarrays into one large sorted array.
The program should prompt the user to input a series of integers. Each goroutine 
which sorts ¼ of the array should print the subarray that it will sort. When 
sorting is complete, the main goroutine should print the entire sorted list. 
**********/

package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "sort"
)

//Read a list of numbers and stores it in a slice.
func read_input() []int{

    read_input := bufio.NewScanner(os.Stdin)
    fmt.Println("\nEnter numbers separated by spaces.")
    fmt.Println("Press enter when done:")
    read_input.Scan()
    input_text := read_input.Text()

    user_input := strings.Split(input_text, " ")

    var user_slice []int

    for i:= 0; i < len(user_input); i++ {
        num, err := strconv.Atoi(user_input[i])
        if err == nil {
            user_slice = append(user_slice, num)
        } else {
            fmt.Println("Invalid input. Please run the program again")
            fmt.Print("Exiting....")
            os.Exit(1)
        }
    }

    return user_slice

}

//decrements remainder till 0
func check_remainder(rem int) (int, int) {

    inc := 0
    if rem > 0 {
        inc = 1
        rem = rem - 1
    }
    return inc, rem
}

//Splits the slice into 4 equal arrays
func split_input(input_slice []int) ([]int, []int, []int, []int){

    length := len(input_slice)
    array_size :=  length / 4
    remainder := length % 4

    var inc1, rem1 = check_remainder(remainder)
    array1 := input_slice[:array_size+inc1]

    var inc2, rem2 = check_remainder(rem1)
    array2 := input_slice[array_size+inc1:array_size*2+inc1+inc2]
    
    var inc3, _ = check_remainder(rem2)
    array3 := input_slice[array_size*2+inc1+inc2:array_size*3+inc1+inc2+inc3]
    
    array4 := input_slice[array_size*3+inc1+inc2+inc3:array_size*4+inc1+inc2+inc3]

    return array1, array2, array3, array4
}

func sort_array(arr []int, c chan[]int){

    orig_arr := make([]int, len(arr))
    copy(orig_arr, arr)
    sort.Ints(arr)
    fmt.Println("Sorted", orig_arr, "into", arr)

    c <- arr

}

func main(){

    input_slice := read_input()

    array1, array2, array3, array4 := split_input(input_slice)

    fmt.Println("\nSorting input ...")

    c := make(chan []int, 4)
    go sort_array(array1, c)
    go sort_array(array2, c)
    go sort_array(array3, c)
    go sort_array(array4, c)
    sort1 := <-c
    sort2 := <-c
    sort3 := <-c
    sort4 := <-c


    var merged []int
    merged = append(merged, sort1...)
    merged = append(merged, sort2...)
    merged = append(merged, sort3...)
    merged = append(merged, sort4...)
    sort.Ints(merged)
    fmt.Println("\nSorted array:")
    fmt.Println(merged)
}
