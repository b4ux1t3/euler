package main

import "fmt"

type card struct{
  rank int
  suit int
}

type hand struct {
  cards []card
  handType int
}

// Hands
const (
  highCard int = iota
  pair
  twoPair
  threeOfAKind
  straight
  flush
  fullHouse
  fourOfAKind
  straightFlush
  royalFlush
  brokenHand
)

// Suits
const (
   clubs int = iota
   diamonds
   hearts
   spades
   brokenSuit
)

// Ranks
const (
  two int = iota
  three
  four
  five
  six
  seven
  eight
  nine
  ten
  jack
  queen
  king
  ace
  brokenRank
)

func main() {
  var cards []card = make([]card, 5)

  cards[0] = card{ten, hearts}
  cards[1] = card{nine, hearts}
  cards[2] = card{three, diamonds}
  cards[3] = card{ace, spades}
  cards[4] = card{queen, clubs}

  var playerOneHand hand = hand{cards, brokenHand}



  var x int = whichHand(playerOneHand)
  fmt.Println(x)
}

func mapHand(input hand) map[int]int {
  var m map[int]int = make(map[int]int)
  for i := 0; i < len(input.cards); i++ {
    // if the
    if val, ok := m[input.cards[i].rank]; ok {
      var _ = val
      m[input.cards[i].rank]++
    } else {
      m[input.cards[i].rank] = 1
    }
  }
  return m
}

// TODO: Determine what hand you have based on a map
func whichHand(input hand) int {
  var m map[int]int = mapHand(input)

  for key, value := range m {
    fmt.Println(key, value)
  }

  return 1
}
// TODO: Sort hands to determine if there is a straight

