package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func swap(nums []int, j int){
	var temp int
	temp = nums[j]
	nums[j]= nums[j+1]
	nums[j+1]=temp
}

func bubblesort(nums []int){

	for i:= 0; i < len(nums) ; i++ {
		for j := 0; j < len(nums)-i-1; j++{
			if nums[j+1] < nums[j]{
				swap(nums,j)
			}
		}
	}

}


func main(){
	fmt.Println("Enter sequence of upto 10 nos(seperate with space).")
	fmt.Println("Press enter when done:\n")
	buf := bufio.NewReader(os.Stdin)
	input,_,_ := buf.ReadLine()
	nos := strings.Split(string(input)," ")
	var nums []int
	for _,s := range(nos){
		no,_ :=strconv.Atoi(s)
		nums = append(nums,no)
	}
	bubblesort(nums)
	fmt.Println(nums)

}

