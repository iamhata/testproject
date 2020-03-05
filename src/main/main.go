package main

import (
  "fmt"
  "github.com/JustinSDK/pkgfoo"
  "test-module/src/objects"
  "test-module/src/util"
  "time"
)

func main() {
  fmt.Println("Hello, world")
  pkgfoo.Hi()
  //add
  var resAdd = util.Add(5,6)
  fmt.Printf("Add: %v ", resAdd)

  var resSubtract = util.Subtract(10, 3)
  fmt.Printf("Substract:%v ", resSubtract)

  var resMultiple  = util.Multiple(5, 4)
  fmt.Printf("Multiple:%v ", resMultiple)

  var resDivide  = util.Divide(2, 3)
  fmt.Printf("Divide:%v ", resDivide)


  goal := make(chan string)

  objects.Test()
  go util.Countdown(goal)
  fmt.Printf(" %s goroutine is finsih!\n", <-goal)



  var resSum = util.Sum(1, 2,3,4,5,6,7,8,9);
  fmt.Printf("Sum:%v ", resSum)

  slice1 := make([]int, 10)
  for i := range slice1 {
    slice1[i] = i
  }

  for _, i := range slice1 {
    fmt.Printf("res:%v \n", i)
  }
  
  time.Sleep(5*time.Second)
}

