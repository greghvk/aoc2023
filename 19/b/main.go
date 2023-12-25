package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var idxPerStr = map[rune]int {
  'x': 0,
  'm': 1,
  'a': 2,
  's': 3,
  ' ': 4, //catchall
}

type op struct {
  idx int
  num uint64
  desc bool
  tgt string
}

var strMap = map[string][]op{}

func main() {
  f := "./19/b/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    if len(scanner.Text()) == 0 {
      break
    }
    parse(scanner.Text())
  }
  res := flow([][]uint64{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}}, "in")
  fmt.Println(res)
}

func parse(s string) {
  els := strings.Split(s, "{")
  tgt := els[0]
  vals := strings.Split(els[1], "}")[0]

  for _, cand := range(strings.Split(vals, ",")) {
    spl := strings.Split(cand, ":")
    if len(spl) == 1 {
      strMap[tgt] = append(strMap[tgt], op{
        idx: 4,
        tgt: spl[0],
      })
      return
    }
    runes := []rune(spl[0])
    desc := runes[1] == '<'
    n := string(runes[2:])
    i, err := strconv.Atoi(n)
    if err != nil {
      panic(err)
    }

    strMap[tgt] = append(strMap[tgt], op{
      idx: idxPerStr[runes[0]],
      num: uint64(i),
      desc: desc,
      tgt: spl[1],
    })

  }
}

func flow(curRanges [][]uint64, cur string) uint64 {
  if cur == "R" || curRanges == nil {
    return 0
  }
  if cur == "A" {
    return prod(curRanges)
  }
  res := uint64(0)

  for i := range strMap[cur] {
    var newRanges [][]uint64
    newRanges, curRanges = splitAndModify(curRanges, strMap[cur][i])
    res += flow(newRanges, strMap[cur][i].tgt)
    if curRanges == nil {
      break
    }
  }
  return res
}

func splitAndModify(r [][]uint64, o op) ([][]uint64, [][]uint64) {
  if o.idx == 4 {
    // return res, nil
    return r, nil
  }
  new := make([][]uint64, len(r))
  cur := make([][]uint64, len(r))

  for i := range(r) {
    new[i] = make([]uint64, len(r[i]))
    cur[i] = make([]uint64, len(r[i]))
    copy(new[i], r[i])
    copy(cur[i], r[i])
  }

  idx := o.idx
  if o.desc {
    if r[idx][0] >= o.num {
      return nil, r
    }
    if r[idx][1] < o.num {
      return new, nil
    }
    new[idx][1] = o.num-1
    cur[idx][0] = o.num
  } else {
    if r[idx][1] <= o.num {
      return nil, r
    }
    if r[idx][0] > o.num {
      return new, nil
    }
    new[idx][0] = o.num+1
    cur[idx][1] = o.num
  }
  return new, cur
}

func prod(ranges [][]uint64) uint64 {
  res := uint64(1)
  for i := 0; i<4; i++ {
    res *= (ranges[i][1] - ranges[i][0] + 1)
  }
  return res
}
