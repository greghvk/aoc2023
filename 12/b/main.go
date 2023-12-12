package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  f := "./12/b/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  res := int64(0)
  for scanner.Scan() {
    str := strings.Fields(scanner.Text())
    inRune := []rune(str[0])
    for i:=0; i<4; i++ {
      inRune = append(inRune, '?')
      inRune = append(inRune, []rune(str[0])...)
    }
    inNums := []int{}
    for i:=0; i<5; i++ {
      inNums = append(inNums, parseNums(str[1])...)
    }

    s := solve(inRune, inNums)
    res += s
  }
  fmt.Println(res)
}

func parseNums(s string) []int {
  els := strings.Split(s, ",")
  var res []int
  for _, el := range(els) {
    i, _ := strconv.Atoi(el)
    res = append(res, i)
  }
  return res
}


func solve(fields []rune, nums []int) int64 {
  dp := make([][]int64, len(fields))
  for i := range(dp) {
    dp[i] = make([]int64, len(nums))
    for j := range(dp[i]) {
      dp[i][j] = -1
    }
  }
  res := doSolve(fields, nums, 0, 0, dp)
  return res
}

func doSolve(fields []rune, nums []int, fieldsPos int, numsPos int, dp [][]int64) int64 {
  if numsPos == len(nums) {
    if allDots(fields, fieldsPos) {
      return 1
    }
    return 0
  }
  if fieldsPos >= len(fields) || numsPos >= len(nums) {
    return 0
  }

  if dp[fieldsPos][numsPos] != -1 {
    return dp[fieldsPos][numsPos]
  }
  switch fields[fieldsPos] {
  case '.':
    dp[fieldsPos][numsPos] = doSolve(fields, nums, fieldsPos+1, numsPos, dp)
    return dp[fieldsPos][numsPos]
  case '#':
    dp[fieldsPos][numsPos] = 0
    if canEnterAtPosition(fields, nums, fieldsPos, numsPos) {
      
      dp[fieldsPos][numsPos] = doSolve(fields, nums, fieldsPos + nums[numsPos] + 1, numsPos+1, dp)
    }
    
    return dp[fieldsPos][numsPos]
  default: // '?'
    cand1 := doSolve(fields, nums, fieldsPos+1, numsPos, dp)
    if canEnterAtPosition(fields, nums, fieldsPos, numsPos) {
      cand1 += doSolve(fields, nums, fieldsPos + nums[numsPos] + 1, numsPos+1, dp)
    }
    dp[fieldsPos][numsPos] = cand1
    return cand1
  }
}


func canEnterAtPosition(fields []rune, nums []int, fieldsPos int, numsPos int) bool {
  if fieldsPos + nums[numsPos] > len(fields) {
    return false
  }
  
  if fieldsPos + nums[numsPos] < len(fields) && fields[fieldsPos + nums[numsPos]] == '#' {
    return false
  }
  for i:=fieldsPos; i<fieldsPos+nums[numsPos]; i++ {
    if fields[i] == '.' {
      return false
    }
  }
  return true

}

func allDots(s []rune, i int) bool {
  for j:=i; j<len(s); j++ {
    if s[j] == '#' {
      return false
    }
  }
  return true
}