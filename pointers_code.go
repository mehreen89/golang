package main

import 
    "fmt"
func main (){

  x := 3
  p := &x             // p, of type *int, points to x
  
  fmt.Println(*p)    //  "3"

  *p = 7            // equivalent to x = 7
  fmt.Println(x, p, *p)   
}
