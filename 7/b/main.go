package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, 2

// 6: Five of a kind, where all five cards have the same label: AAAAA
// 5: Four of a kind, where four cards have the same label and one card has a different label: AA8AA
// 4: Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
// 3: Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// 2: Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// 1: One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// 0: High card, where all cards' labels are distinct: 23456


var mapping = map[rune]int {
  'A': 13,
  'K': 12,
  'Q': 11,
  'T': 10,
  '9': 8,
  '8': 7,
  '7': 6,
  '6': 5,
  '5': 4,
  '4': 3,
  '3': 2,
  '2': 1,
  'J': 0,
  
}
func getType(s string) int {
  maxCnt := 0
  numPairs := 0
  numJokers := 0
  m := make(map[rune]int)
  for _, r := range(s) {
    if r == 'J' {
      numJokers++
      continue
    }
    m[r]++
    if m[r] == 2 {
      numPairs++
    }
    if m[r] > maxCnt {
      maxCnt = m[r]
    }
  }
  // try 4 or 5
  if maxCnt + numJokers > 3 {
    return maxCnt+numJokers+1 // 4 or 5
  }
  // try full house
  if maxCnt+numJokers == 3 && numPairs == 2 {
    return 4
  }
  // try 3
  if maxCnt + numJokers == 3 {
    return 3
  }
  //try 2 pairs
  if numPairs == 2 || numPairs == 1 && numJokers ==1 {
    return 2
  }
  // try 1 pair
  if numJokers == 1 || numPairs == 1 {
    return 1
  }
  return 0
  

}

func compHands(s1, s2 []rune) bool {
  for i := range(s1) {
    if s1[i] == s2[i] {
      continue
    }
    return mapping[s1[i]] < mapping[s2[i]]
  }
  return true
}

type state struct {
  t int
  hand string
  bet int
}

func main() {
  f := "./7/b/input"
  m := make(map[string]int)
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  var els []state

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    t := strings.Split(scanner.Text(), " ")
    i, err := strconv.Atoi(t[1])
    if err != nil {
      panic(err)
    }
    els = append(els, calcState(t[0], i))
    m[t[1]] = i
  }
  sort.Slice(els, func(i, j int) bool {
    if els[i].t != els[j].t {
      return els[i].t < els[j].t
    }
    return compHands([]rune(els[i].hand), []rune(els[j].hand))
  })
  var res int64
  for i := range(els) {
    // fmt.Println("adding ", i+1, " times ", els[i].bet)
    // fmt.Println(els[i])
    if i > 0 && els[i-1]==els[i] && !compHands([]rune(els[i-1].hand), []rune(els[i].hand)) {
      panic(string("incorrect: " + els[i-1].hand +  " "+ els[i].hand))
    }
    res += int64(i+1) * int64(els[i].bet)
  }
  fmt.Println(res)
}

func calcState(hand string, bet int) state {
  return state{
    t: getType(hand),
    hand: hand,
    bet: bet,
  }
}