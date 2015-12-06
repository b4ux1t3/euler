package main

import "fmt"

type card struct{
  rank int
  suit int
}

const (
   clubs int = iota
   diamonds
   hearts
   spades
   brokenSuit
)

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
  var test card = parseCard("TH")
  fmt.Println(test.rank, test.suit)
}

func parseCard(inputCard string) card {
  if len(inputCard) != 2{
    return card{brokenRank, brokenSuit}
  } else {
    var newCard card = card{brokenRank, brokenSuit}
    var rank string = string(inputCard[0])
    var suit string = string(inputCard[1])

    switch rank {
    case "2":
      newCard.rank = two
    case "3":
      newCard.rank = three
    case "4":
      newCard.rank = four
    case "5":
      newCard.rank = five
    case "6":
      newCard.rank = six
    case "7":
      newCard.rank = seven
    case "8":
      newCard.rank = eight
    case "9":
      newCard.rank = nine
    case "T":
      newCard.rank = ten
    case "J":
      newCard.rank = jack
    case "Q":
      newCard.rank = queen
    case "K":
      newCard.rank = king
    case "A":
      newCard.rank = ace
    default:
      newCard.rank = brokenRank
    }

    switch suit {
    case "C":
      newCard.suit = clubs
    case "D":
      newCard.suit = diamonds
    case "H":
      newCard.suit = hearts
    case "S":
      newCard.suit = spades
    default:
      newCard.suit = brokenSuit
    }

    return newCard
  }
}

