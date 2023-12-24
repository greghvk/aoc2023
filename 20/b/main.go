package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// We need to have low signal coming from gates listed below at once.
// Why? rx (our target) only receives signals from 1 AND gate, which in turn receives from 4 AND gates.
// Befre that, these 4 AND only receive signals from other 4 AND gates listed below (see gates.drawio for visualisation).
// It turns out that the listed gates (nf, pm, jd, qm) are periodically sending a low signal (run and see logs).
// The requency of each of these gates is a prime. The moment where all of the signals are low
// is the LCM of these primes (which is just multiply):
// Result: 3917 * 4057 * 3943 * 3931 =
var toLog = map[string]bool {
  "nf": true,
  "pm": true,
  "jd": true,
  "qm": true,
}

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
  
  for true {
    cnt++
    if process(cnt) {
      break
    }
    if cnt == 10000 {
      break
    }
  }

  fmt.Println(cnt)
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

func process(press int) bool {

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
    if toLog[sign.source] && !sign.high {
      fmt.Println("sending high from", sign.source, "at", press)
    }
    // pm: 3917
    // jd: 4057
    // nf: 3943
    // qm: 3931
    if sign.target == "rx" && !sign.high {
      return true
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

  return false
}