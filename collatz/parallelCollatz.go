package main

import "fmt"

func main(){
  //start the Goroutine for computing the upper half of our range
  go upperHalf()

  //start the lower half
  lowerHalf()
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

func lowerHalf(){
  for j := 0; j < 500000; j++{
    collatz(j + 1)
  }
}

func upperHalf(){
  for i := 500000; i < 1000000; i++{
    collatz(i + 1)
  }
}

