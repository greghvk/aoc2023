package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  f := "./9/a/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  var res int64
  for scanner.Scan() {
    // fmt.Println("solving ", scanner.Text())
    f := getFields(scanner.Text())
    x := lastSeqValue(f) + int64(f[len(f)-1])
    // fmt.Println("adding ", x)
    res += x
  }
  fmt.Println(res)
}

func lastSeqValue(els []int) int64 {
  onlyZero := true
  seq := make([]int, len(els)-1)
  for i := 1; i < len(els); i++ {
    seq[i-1] = els[i] - els[i-1]
    if seq[i-1] != 0 {
      onlyZero = false
    }
  }
  // fmt.Println("seq is ", seq)
  if onlyZero {
    return 0
  }
  return int64(seq[len(seq)-1]) + lastSeqValue(seq)
}

func getFields(s string) []int {
  var res []int
  for _, el := range(strings.Fields(s)) {
    i, _ := strconv.Atoi(el)
    res = append(res, i)
  }
  return res
}