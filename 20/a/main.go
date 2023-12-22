package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type element struct {
  statePerInput map[string]bool
  t int
  outs []string
  state bool
}

var m = map[string]element{}

var typePerChar = map[rune]int{
  '%': 1,
  '&': 2,
}

func main() {
  f := "./20/a/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    parse(scanner.Text())
  }
  fmt.Println(m)
  cnt := 0
  var resLow uint64
  var resHigh uint64
  
  for cnt < 1000 {
    cnt++
    nLow, nHigh := process()
    resLow += nLow
    resHigh += nHigh
  }
  fmt.Println(resLow * resHigh)
}

func parse(s string) {
  spl := strings.Split(s, " -> ")
  in := []rune(spl[0])
  inType := typePerChar[in[0]]
  inName := string(in[1:])
  if spl[0] == "broadcaster" {
    inType = 0
    inName = spl[0]
  }
  
  
  var outs []string
  for _, out := range(strings.Split(spl[1], ", ")) {
    outs = append(outs, out)
    el := m[out]
    if el.statePerInput == nil {
      el.statePerInput = make(map[string]bool)
    }
    el.statePerInput[inName] = false
    m[out] = el
  }
  prev := m[inName]
  prev.t = inType
  prev.outs = outs
  m[inName] = prev
}

type signal struct {
  high bool
  target, source string
}

func process() (uint64, uint64) {
  var nLow uint64
  var nHigh uint64
  nLow++

  q := []signal{}
  for _, tgt := range(m["broadcaster"].outs) {
    q = append(q, signal{
      high: false,
      target: tgt,
      source: "broadcaster",
    })
  }

  for len(q) > 0 {
    sign := q[0]
    q = q[1:]
    if sign.high {
      nHigh++
    } else {
      nLow++
    }
    switch m[sign.target].t {
    case 0: // sink
      continue
    case 1: // % flip flop
      el := m[sign.target]
      if sign.high {
        continue
      }
      el.state = !el.state
      for _, tgt := range el.outs {
        q = append(q, signal{
          high: el.state,
          source: sign.target,
          target: tgt,
        })
      }
      m[sign.target] = el
    case 2: // & 
      el := m[sign.target]
      el.statePerInput[sign.source] = sign.high
      resHigh := false
      for _, v := range(el.statePerInput) {
        if !v {
          resHigh = true
          break
        }
      }
      for _, tgt := range el.outs {
        q = append(q, signal{
          high: resHigh,
          source: sign.target,
          target: tgt,
        })
      }
      

    }

  }

  return nLow, nHigh
}