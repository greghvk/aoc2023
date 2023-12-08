package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
  f := "./5/a/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()


  scanner := bufio.NewScanner(file)

  scanner.Scan()
  els := strings.Split(scanner.Text(), ": ")
  fmt.Println(els)
  nums:= parseStrings(strings.Split(els[1], " "))
  fmt.Println(nums)


  scanner.Scan()
  for {
    mapping := getMapping(scanner)
    if len(mapping) == 0 {
      break
    }
    nums = applyMapping(nums, mapping)
  }
  fmt.Println("final nums ", nums)
  min := nums[0]
  for _, el := range(nums) {
    if el < min {
      min = el
    }
  }
  fmt.Println(min)
}

func parseStrings(s []string) []int64 {
  var res []int64
  for _, cur := range(s) {
    res = append(res, mustParse(cur))
  }
  return res
}

func mustParse(s string) int64 {
  res, err := strconv.ParseInt(s, 10, 64)
  if err != nil {
    panic(err)
  }
  return res
}

func getMapping(scanner *bufio.Scanner) [][]int64 {
  var res [][]int64 
  if !scanner.Scan() {
    return res
  }
  for scanner.Scan() && len(scanner.Text()) > 0 {
    els := parseStrings(strings.Split(scanner.Text(), " "))
    res = append(res, []int64{els[1], els[0], els[2]})
  }
  sort.Slice(res, func(i, j int) bool {
    return res[i][0] < res[j][0]
  })
  return res
}

func applyMapping(srcs []int64, mapping [][]int64) []int64 {
  sort.Slice(srcs, func(i, j int) bool {
    return srcs[i] < srcs[j]
  })
  var res []int64
  srcPos := 0
  mapPos := 0
  // fmt.Println("using mapping ", mapping)

  for srcPos < len(srcs) {
    for mapPos < len(mapping) && mapping[mapPos][0] + mapping[mapPos][2] < srcs[srcPos] {
      mapPos++
    }

    if mapPos < len(mapping) && mapping[mapPos][0] <= srcs[srcPos] {
      diff := srcs[srcPos] - mapping[mapPos][0]
      res = append(res, mapping[mapPos][1] + diff)
      // fmt.Println("mapping ", srcs[srcPos], " to ", mapping[mapPos][1] + diff)
    } else {
      // fmt.Println("mapping ", srcs[srcPos], " to itself")
      res = append(res, srcs[srcPos])
    }
    srcPos++
  }
  return res
}