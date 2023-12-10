package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
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
  'A': 5,
  'K': 4,
  'Q': 3,
  'J': 2,
  'T': 1,
  
}
func getType(s string) int {
  maxCnt := 0
  numPairs := 0
  m := make(map[rune]int)
  for _, r := range(s) {
    m[r]++
    if m[r] == 2 {
      numPairs++
    }
    if m[r] > maxCnt {
      maxCnt = m[r]
    }
  }
  if maxCnt >3 {
    return maxCnt+1
  }
  if maxCnt == 3{
    if numPairs == 2 {
      return 4
    }
    return 3
  }
  
  if numPairs == 2 {
    return 2
  }
  if maxCnt == 2 {
    return 1
  }
  return 0

}

func compHands(s1, s2 []rune) bool {
  for i := range(s1) {
    if s1[i] == s2[i] {
      continue
    }
    if unicode.IsDigit(s1[i]) || unicode.IsDigit(s2[i]) {
      return s1[i] < s2[i]
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
  f := "./7/a/input"
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