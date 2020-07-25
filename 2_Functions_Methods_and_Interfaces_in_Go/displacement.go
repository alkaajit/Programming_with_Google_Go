//displacement.go
package main

import "fmt"

func main() {
	
  var a float64
  fmt.Print("Enter acceleration:")
  fmt.Scanln(&a)

  var v0 float64
  fmt.Print("Enter the initial velocity:")
  fmt.Scanln(&v0)
  
  var s0 float64
  fmt.Print("Enter initial displacement:")
  fmt.Scanln(&s0)
  
  var t float64
  fmt.Print("Enter time:")
  fmt.Scanln(&t)

  fn := GenDisplaceFn(a, v0, s0)
  fmt.Println("Displacement after", t ,"seconds:",fn(t))


}

func GenDisplaceFn(a,v0,s0 float64) func(float64) float64{
	return func(t float64) float64 {
		return 0.5*a*t*t+v0*t+s0
	}
}
