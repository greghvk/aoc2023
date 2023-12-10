package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  f := "./6/b/input_test"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Scan()
  time := getField(scanner.Text())
  scanner.Scan()
  distance := getField(scanner.Text())

  fmt.Println("res is ", solve(time, distance))
}

func getField(s string) int64 {
  var concat string
  for _, el := range(strings.Fields(s)[1:]) {
    concat += el
  }
  i, err := strconv.ParseInt(concat, 10, 64)
  if err != nil {
    panic(err)
  }
  return i
}

func solve(time, dist int64) int64 {
  l := int64(0)
  r := time+1
  min:= findMin(l, r, time, dist)
  max := findMax(min, r, time, dist)
  fmt.Println("min ", min, " max ", max)
  return max-min
}

func findMin(l, r, time, dist int64) int64 {
  for l < r {
    m := (l+r)/2
    d1 := getDist(m, time)
    d2 := getDist(m+1, time)
    if d2 < d1 {
      r = m-1
      continue
    }
    if d1 > dist {
      r = m-1
    } else {
      l = m+1
    }
  }
  return l
}

func getDist(chargeTime, time int64) int64 {
  timeLeft := time-chargeTime
  return chargeTime * timeLeft
}

func findMax(l, r, time, dist int64) int64 {
  for l < r {
    m := (l+r)/2
    d1 := getDist(m, time)
    d2 := getDist(m+1, time)
    if d1 < d2 {
      l = m+1
      continue
    }
    if d1 >= dist {
      l = m+1
    } else {
      r = m-1
    }
  }
  return l
}