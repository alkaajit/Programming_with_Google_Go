//read.go

package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

type name struct {
    fname string
    lname string
}

func main(){

    input := bufio.NewScanner(os.Stdin)
    fmt.Printf("\nPlease enter the filename containing a list of names: ")
    input.Scan()
    filename := input.Text()

    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("\nFile could not be opened. Please make sure you have the")
        fmt.Println("correct filename and path before running again.")
        os.Exit(1)	
    }

    name_slice := make([]name, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        name_strings := strings.Fields(scanner.Text())
        output := name{fname: name_strings[0], lname: name_strings[1]}
        name_slice = append(name_slice, output)
    }

    fmt.Println("\nHere are your list of names from", filename)

    for i:=0; i<len(name_slice); i++{
	    fmt.Println("First name: " + name_slice[i].fname, ", Last name: " + name_slice[i].lname)
    }

}

