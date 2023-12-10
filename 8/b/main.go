package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type solver struct {
  m map[string][]string
  curNodes []string
}

func (s *solver) addEdges(str string) {
  els := strings.Split(str, " = ")
  edges := strings.Split(els[1], ", ") 
  s.m[els[0]] = append(s.m[els[0]], edges[0][1:], edges[1][:len(edges[1])-1])
  if els[0][len(els[0])-1] == 'A' {
    s.curNodes = append(s.curNodes, els[0])
  }
}

func (s *solver) numAtZ() int {
  res := 0
  for i := range(s.curNodes) {
    if s.curNodes[i][len(s.curNodes[i])-1] == 'Z' {
      res++
    }
  }
  return res
}

func main() {
  f := "./8/b/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Scan()
  dirs := []rune(scanner.Text())
  scanner.Scan()
  s := &solver{
    m: make(map[string][]string),
  }
  for scanner.Scan() {
    s.addEdges(scanner.Text())
  }
  fmt.Println(len(dirs))
  var results []int

  for i := range(s.curNodes) {
    pos := 0
    steps := 0
    for {
      
        if dirs[pos] == 'L' {
          s.curNodes[i] = s.m[s.curNodes[i]][0]
        } else {
          s.curNodes[i] = s.m[s.curNodes[i]][1]
        }
        pos++
        if pos == len(dirs) {
          pos = 0
        }
        steps++
        if s.curNodes[i][len(s.curNodes[i])-1] == 'Z'{
          break
        }
      }

      results = append(results, steps)
      
      // if steps == 100 {
      //   break
      // }
    
  }
  res := int64(results[0])
  for i, r := range(results) {
    if i == 0 {
      continue
    }
    res = int64(r) * res / 307
  }
  fmt.Println(res)

}