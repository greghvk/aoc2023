package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// one two six ten - len 3
// four five nine - len 4
// three seven eight - len 5

func main() {
  m := map[string]int{
    "zero": 0,
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
  }
  f := "./1/input"
  file, err := os.Open(f)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    res := 0
    for scanner.Scan() {
      cur := calculate(scanner.Text(), m)
      fmt.Println(cur)
      res += cur
    }
    fmt.Println("sum is ", res)
  
}


func calculate(s string, m map[string]int) int {
  first := -1
  last := -1
  runes := []rune(s)
  for i, char := range(runes) {
    var dig int
    var ok bool
    if !unicode.IsDigit(char) {
      dig, ok = tryGetDig(i, runes, m)
      if !ok {
        continue
      }
    } else {
      dig = int(char-'0')
    }

    if first == -1 {
      first = dig
    }
    last = dig
  }
  return 10*first + last
}

func tryGetDig(i int, chars []rune, m map[string]int) (int, bool) {
  var cur []rune
  for j := i; j>i-5; j-- {
    if j < 0 {
      break
    }
    cur = append([]rune{chars[j]}, cur...)
    if v, ok := m[string(cur)]; ok {
      return v, true
    }
  }
  return -1, false
}