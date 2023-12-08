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
  f := "./5/b/input_test"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()


  scanner := bufio.NewScanner(file)

  scanner.Scan()
  els := strings.Split(scanner.Text(), ": ")
  nums:= getNums(els[1])


  scanner.Scan()
  for {
    mapping := getMapping(scanner)
    if len(mapping) == 0 {
      break
    }
    nums = applyMapping(nums, mapping)
  }
  min := nums[0][0]
  for _, el := range(nums) {
    if el[0] < min {
      min = el[0]
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

func applyMapping(srcs [][]int64, mapping [][]int64) [][]int64 {
  sort.Slice(srcs, func(i, j int) bool {
    return srcs[i][0] < srcs[j][0]
  })
  var res [][]int64
  srcPos := 0
  mapPos := 0

  for srcPos < len(srcs) {
    for mapPos < len(mapping) && mapping[mapPos][0] + mapping[mapPos][2] < srcs[srcPos][0] {
      mapPos++
    }

    l := srcs[srcPos][0]
    r := l + srcs[srcPos][1]

    for l < r {
      if mapPos == len(mapping) {
        res = append(res, []int64{l, r-l})
        l = r
        break
      }
      if mapping[mapPos][0] > l {
        dist := mapping[mapPos][0] - l
        if l+dist > r {
          dist = r-l
        }
        res = append(res, []int64{l, dist})
        l += dist
        continue
      }
      if mapping[mapPos][0] + mapping[mapPos][2] <= l {
        mapPos++
        continue
      }
      dist := mapping[mapPos][0] + mapping[mapPos][2] - l
      if l+dist > r {
        dist = r-l
      }
      diff := mapping[mapPos][1] - mapping[mapPos][0]
      res = append(res, []int64{l+diff, dist})
      l += dist
    }
    
    srcPos++
  }
  return res
}

func getNums(s string) [][]int64 {
  normal := parseStrings(strings.Split(s, " "))
  var res [][]int64
  for i := 0; i<len(normal); i++ {
    res = append(res, []int64{normal[i], normal[i+1]})
    i++
  }
  return res
}