package main

import (
  "fmt"
  "os"
  "strings"
)

type card struct{
  rank int
  suit int
}

type hand struct {
  firstCard card
  secondCard card
  thirdCard card
  fourthCard card
  fifthCard card
  handType int
}

// Hands
const(
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
  // Open our poker hands file, checks for an error
  f, err := os.Open("poker.txt")
  check(err)

  var test string = getLine(f, 0)

  playerOneHand, playerTwoHand := parseHands(parseCards(test))

  fmt.Println(test)
  fmt.Println("Player One's Hand:", playerOneHand)
  fmt.Println("Player Two's Hand:", playerTwoHand)
}

// Reads a single line of the poker hands file
func getLine(f *os.File, offset int64) string{
  // Array storing how many buytes we want to read
  bytesToRead := make([]byte, 30)

  // Seek to the beginning of the line we want to read (0 indexed)
  seek, err := f.Seek(offset * 30, 0)
  check(err)
  fmt.Println("Line:", seek / 30)

  // Store that line in bytesToRead
  readBytes, err := f.Read(bytesToRead)
  check(err)
  fmt.Println("Cards read:", readBytes / 3)

  // Change that to a string
  var line string = string(bytesToRead)

  return line
}

// Parses a line into an array of cards
func parseCards(line string) [10]card{
  var cardStrings []string = strings.Fields(line)

  var cards [10]card

  for i := 0; i < len(cardStrings); i++{
    cards[i] = parseCard(cardStrings[i])
  }

  return cards
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

// Parses an array of cards into two hands, returns two empty arrays if there
// are not exactly ten cards in the input
// [0] [1] [2] [3] [4] [5] [6] [7] [8] [9]
func parseHands(cards [10]card) (hand, hand){
  if len(cards) != 10 {
    return hand{}, hand{}
  } else {
    var playerOne []card = cards[:5]
    var playerTwo []card = cards[5:]

    var playerOneHand hand = hand{playerOne[0],
                                  playerOne[1],
                                  playerOne[2],
                                  playerOne[3],
                                  playerOne[4],
                                  brokenHand}
    var playerTwoHand hand = hand{playerTwo[0],
                                  playerTwo[1],
                                  playerTwo[2],
                                  playerTwo[3],
                                  playerTwo[4],
                                  brokenHand}

    return playerOneHand, playerTwoHand
  }
}

func mapHand(input hand) map[int]int {
  var m map[int]int = make(map[int]int)
  for i := 0; i < len(input.cards); i++ {
    if val, ok := m[input.cards[i].rank]; ok {
      var _ = val
      m[input.cards[i].rank]++
    } else {
      m[input.cards[i].rank] = 1
    }
  }
  return m
}

// Error checking
func check(e error) {
    if e != nil {
        panic(e)
    }
}

