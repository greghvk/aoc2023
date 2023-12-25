package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)


var vPerDir = map[int][]int64{
  0: {0, 1},
  1: {1, 0},
  2: {0, -1},
  3: {-1, 0},
}

func main() {
  f := "./18/a/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()
  points := [][]int64{{0, 0}}
  pos := []int64{0, 0}

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    l := strings.Split(scanner.Text(), "#")[1]
    l = strings.Split(l, ")")[0]
    runes := []rune(l)
    dir, err := strconv.Atoi(string(runes[len(runes)-1]))
    if err != nil {
      panic(err)
    }
    i, err := strconv.ParseInt(string(runes[:len(runes)-1]), 16, 32)
    if err != nil {
      panic(err)
    }
    v := vPerDir[dir]
    
    pos[0] += v[0] * i
    pos[1] += v[1] * i
    p2 := []int64{pos[0], pos[1]}
    if pos[0] == 0 && pos[1] == 0 {
      break
    }
    points = append(points, p2)
  }

  sort.Slice(points, func(i, j int) bool {
    if points[i][0] == points[j][0] {
      if points[i][1] == points[j][1] {
        panic(fmt.Sprintf("duplicate point %v %v", points[i], points[j]))
      }
      return points[i][1] < points[j][1]
    }
    return points[i][0] < points[j][0]
  })

  fmt.Println(solve(points))
}

var prevH int64
func solve(points [][]int64) int64 {
  i := 0
  prevRow := []int64{}
  var res int64

  for i < len(points) {
    curH := points[i][0]
    curRow := []int64{}
    for i < len(points) && points[i][0] == curH {
      curRow = append(curRow, points[i][1])
      i++
    }

    res += (curH - prevH - 1) * sumDistances(prevRow)
    merged, toAdd := mergeAndCount(prevRow, curRow) 
    prevRow = merged
    res += toAdd
    prevH = curH

  }
  return res

}

func sumDistances(row []int64) int64 {
  if len(row) % 2 == 1 {
    panic(fmt.Sprintf("uneven  row %v", row))
  }
  var res int64
  for i := 1; i<len(row); i = i+2 {
    res += row[i] - row[i-1] + 1
  }
  return res
}

func mergeAndCount(prevRow, curRow []int64) (mergedRows []int64, res int64) {
  prevCol := int64(-1)
  curState := 0
  prevPos := 0
  curPos := 0

  for prevPos < len(prevRow) || curPos < len(curRow) {
    res++
    if curPos == len(curRow) || (prevPos < len(prevRow) && prevRow[prevPos] < curRow[curPos]) { // 1 prev
      if curState != 0 {
        res += prevRow[prevPos] - prevCol - 1
      }
      curState = transition(curState, true, false)

      mergedRows = append(mergedRows, prevRow[prevPos])
      prevCol = prevRow[prevPos]
      prevPos++
    } else if prevPos == len(prevRow) || (curPos < len(curRow) &&  prevRow[prevPos] > curRow[curPos]) { // 1 cur
      if curState != 0 {
        res += curRow[curPos] - prevCol - 1
      }
      curState = transition(curState, false, true)

      mergedRows = append(mergedRows, curRow[curPos])
      prevCol = curRow[curPos]
      curPos++
    } else { // 2 points
      if curRow[curPos] != prevRow[prevPos] {
        panic("what?")
      }
      if curState != 0 {
        res += prevRow[prevPos] - prevCol - 1
      }
      curState = transition(curState, true, true)

      
      prevCol = curRow[curPos]
      prevPos++
      curPos++
    }
  }
  return mergedRows, res
}

func transition(curState int, hasPrev, hasCur bool) int {
  switch curState{
  case 0:
    // outside - 0
      // 1 point prev -> inside, add
      // 1 point cur -> only below, add
      // 2 points -> only above, skip
    if hasPrev && hasCur {
      return 2
    } else if hasPrev  {
      return 3
    } else if hasCur {
      return 1
    }
  case 1:
    // only below - 1
      // 1 point prev -> painc
      // 1 point cur -> outside, add
      // 2 points -> inside, skip
    if hasPrev && hasCur {
      return 3
    } else if hasPrev {
      panic("below has prev")
    } else if hasCur {
      return 0
    }
  case 2:
    // only above - 2
      // 1 point prev -> panic
      // 1 point cur -> inside, add
      // 2 points -> outside, skip
    if hasPrev && hasCur {
      return 0
    } else if hasPrev {
      panic("above has prev")
    } else if hasCur {
      return 3
    }
  case 3:
    // fully inside - 3
      // 1 point prev -> outside, add
      // 1 point cur -> above, add
      // 2 points -> only below, skip 
    if hasPrev && hasCur {
      return 1
    } else if hasPrev {
      return 0
    } else if hasCur {
      return 2
    }
  default:
    panic("wrong state")
  }
  panic("unreachable")
}