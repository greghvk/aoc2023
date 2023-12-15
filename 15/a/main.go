package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type element struct {
  label string
  focal int
}

type hashMap struct {
  m map[int][]element
}

func main() {
  f := "./15/b/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  m := hashMap{
    m: make(map[int][]element),
  }
  scanner := bufio.NewScanner(file)
  scanner.Scan()
  txt := strings.Split(scanner.Text(), ",")
  
  for i := range(txt) {
    m.process(txt[i])
  }
  res := 0
  for k, v := range m.m {
    for i := range v {
      res += (k+1) * (i+1) * v[i].focal
    }
  }
  fmt.Println(res)
}


func (m* hashMap) process(s string) {
  txt := strings.Split(s, "=")
  if len(txt) == 1 {
    els := []rune(s)
    els = els[:len(els)-1]
    m.remove(string(els))
  } else {
    i, _ :=strconv.Atoi(txt[1])
    m.add(txt[0], i)
  }
}

func (m *hashMap) remove(s string) {
  h := hash([]rune(s))
  for i := range(m.m[h]) {
    if m.m[h][i].label == s {
      m.m[h] = append(m.m[h][:i], m.m[h][i+1:]...)
      return
    }
  }
}

func (m hashMap) add(s string, pow int) {
  h := hash([]rune(s))
  for i := range(m.m[h]) {
    if m.m[h][i].label == s {
      m.m[h][i].focal = pow
      return
    }
  }
  m.m[h] = append(m.m[h], element{
    label: s,
    focal: pow,
  })
  
}

func hash(r []rune) int {
  res := 0
  for _, c := range r {
    res += int(c)
    res *= 17
    res %= 256
  }
  return res
}