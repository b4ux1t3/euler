package main

import "fmt"

func main(){
  var highest int = 0
  var highestInput = 0

  var result int = 0


  for i := 0; i < 1000000; i++{
    //fmt.Println("Collatzing: ", i + 1)
    result = collatz(i + 1)

    if result > highest{
      highest = result
      highestInput = i + 1
    }
  }

  fmt.Println("Most steps:", highestInput, "," , highest)
}

// Takes an input. Returns the number of steps it takes to get to 1 via
// the Collatz Conjecture
func collatz(input int) int{
  var count int = 0

  fmt.Println("Collatzing:", input)
  
  for input > 1 {
    // If number is even, divide by 2, else multiply by 3 and add 1
    if input % 2 == 0 {
      input /= 2
      count++
    } else {
      input = (input * 3) + 1
      count++
    }
  }
  // Include the number 1
  count++
  return count
}

