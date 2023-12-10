package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type solver struct {
  m map[string][]string
}

func (s *solver) addEdges(str string) {
  els := strings.Split(str, " = ")
  edges := strings.Split(els[1], ", ") 
  s.m[els[0]] = append(s.m[els[0]], edges[0][1:], edges[1][:len(edges[1])-1])
}

func main() {
  f := "./8/a/input"
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

  cur := "AAA"
  pos := 0
  steps := 0
  for cur != "ZZZ" {
    if dirs[pos] == 'L' {
      cur = s.m[cur][0]
    } else {
      cur = s.m[cur][1]
    }
    pos++
    if pos == len(dirs) {
      pos = 0
    }
    steps++
  }
  fmt.Println(steps)

}