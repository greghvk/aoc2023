package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var dirs = [][]int {
  {-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

var symbols = map[rune][][]int{
  '|': {{1, 0}, {-1, 0}},
  '-': {{0, -1}, {0, 1}},
  'L': {{-1, 0}, {0, 1}},
  'J': {{-1, 0}, {0, -1}},
  '7': {{1, 0}, {0, -1}},
  'F': {{1, 0}, {0, 1}},
}

func main() {
  f := "./10/b/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  var graph [][]rune
  scanner := bufio.NewScanner(file)
  pos := [][]int{}
  for scanner.Scan() {
    graph = append(graph, []rune(scanner.Text()))
    for i, el := range(graph[len(graph)-1]) {
      if el == 'S' {
        pos = append(pos, []int{len(graph)-1, i})
      }
    }
  }

  replaceS(graph, pos[0])
  

  visited := make(map[string]bool)
  visited[makeKey(pos[0])] = true
  res := 0
  
  for len(pos) > 0 {
    sz := len(pos)
    for ; sz > 0 ; sz-- {
      cur := pos[0]
      i := cur[0]
      j := cur[1]
      pos = pos[1:]
      for _, dir := range(symbols[graph[i][j]]) {
        newi := i+dir[0]
        newj := j+dir[1]
        key := makeKey([]int{newi, newj})
        if visited[key] {
          continue
        }
        
        if !canTravel(newi, newj, graph, -dir[0], -dir[1]) {
          continue
        }
        visited[key] = true
        pos = append(pos, []int{newi, newj})
      }

    }
  }
  newg, newVisited := transformGraph(graph, visited)
  for i := 0; i<len(newg); i++ {
    dfs(newg, i, 0, newVisited)
    dfs(newg, i, len(newg[0])-1, newVisited)
  }
  for j := 0; j<len(newg[0]); j++ {
    dfs(newg, 0, j, newVisited)
    dfs(newg, len(newg)-1, j, newVisited)
  }
  for i := range(newg) {
    for j := range(newg[0]) {
      if !newVisited[makeKey([]int{i, j})] && newg[i][j] != ',' && newg[i][j] != 'x' {
        res++
      }
    }
  }
  fmt.Println(res)
}

func makeKey(pos []int) string {
  return strconv.Itoa(pos[0]) + "," + strconv.Itoa(pos[1])
}

func canTravel(i, j int, graph [][]rune, i_dir, j_dir int) bool {
  if i<0 || i == len(graph) || j < 0 || j == len(graph[0]) {
    return false
  }
  if graph[i][j] == '.' {
    return false
  }
  return sliceHas(symbols[graph[i][j]], []int{i_dir, j_dir})
}

func sliceHas(s [][]int, target []int) bool {
  for _, t := range(s) {
    if t[0] == target[0] && t[1] == target[1] {
      return true
    }
  }
  return false
}


func transformGraph(g [][]rune, visited map[string]bool) ([][]rune, map[string]bool) {
  res := make([][]rune, 2*len(g)-1)
  resVisited := make(map[string]bool)
  for i := range(res) {
    res[i] = make([]rune, 2*len(g[0])-1)
  }
  for i := 0; i<len(g); i++ {
    for j:=0; j<len(g[0]); j++ {
      res[2*i][2*j] = g[i][j]
      k := makeKey([]int{i, j})
      if visited[k] {
        resVisited[makeKey([]int{2*i, 2*j})] = true
      }
      if i < len(g)-1 {
        k2 := makeKey([]int{i+1, j})
        if verticallyConnected(g[i][j], g[i+1][j]) && visited[k] && visited[k2] {
          resVisited[makeKey([]int{2*i+1, 2*j})] = true
          res[2*i + 1][2*j] = '|'
        } else {
          res[2*i + 1][2*j] = ','
        }
      }
      if j < len(g[0])-1 {
        k2 := makeKey([]int{i, j+1})
        if horizontallyConnected(g[i][j], g[i][j+1]) && visited[k] && visited[k2] {
          resVisited[makeKey([]int{2*i, 2*j+1})] = true
          res[2*i][2*j+1] = '-'
        } else {
          res[2*i][2*j+1] = ','
        }
      }
      if i < len(g)-1 && j < len(g[0])-1 {
        res[2*i+1][2*j+1] = ','
      }
    }
  }
  return res, resVisited
}

var toRight = []rune{
  '-', 'L', 'F',
}

var fromLeft = []rune{
  '-', '7', 'J',
}

var toDown = []rune{
  '|', '7', 'F',
}

var fromUp = []rune{
  '|', 'L', 'J',
}

func verticallyConnected(u, d rune) bool {
  return contains(toDown, u) && contains(fromUp, d)
}
func horizontallyConnected(l, r rune) bool {
  return contains(toRight, l) && contains(fromLeft, r)

}
func contains(s []rune, r rune) bool {
  for _, i := range(s) {
    if r == i {
      return true
    }
  }
  return false
}

func dfs(g [][]rune, i, j int, visited map[string]bool) {
  if visited[makeKey([]int{i, j})] || g[i][j] == 'x' {
    return
  }
  g[i][j] = 'x'
  for _, dir := range dirs {
    newi := i+dir[0]
    newj := j+dir[1]
    if newi < 0 || newi == len(g) || newj < 0 || newj == len(g[0]) {
      continue
    }
    dfs(g, newi, newj, visited)
  }
}

func replaceS(graph [][]rune, pos []int){
  i := pos[0]
  j := pos[1]
  
  if i > 0 && i < len(graph)-1 && contains(toDown, graph[i-1][j]) && contains(fromUp, graph[i+1][j]) {
    graph[i][j] = '|'
  } else if j>0 && j<len(graph[0])-1 && contains(toRight, graph[i][j-1]) && contains(fromLeft, graph[i][j+1]) {
    graph[i][j] = '-'
  } else if i>0 && j<len(graph[0])-1 && contains(toDown, graph[i-1][j]) && contains(fromLeft, graph[i][j+1]) {
    graph[i][j] = 'L'
  } else if i>0 && j>0 && contains(toDown, graph[i-1][j]) && contains(toRight, graph[i][j-1]) {
    graph[i][j] = 'J'
  } else if i<len(graph)-1 && j>0 && contains(fromUp, graph[i+1][j]) && contains(toRight, graph[i][j-1]) {
    graph[i][j] = '7'
  } else if i<len(graph)-1 && j<len(graph[0])-1 && contains(fromUp, graph[i+1][j]) && contains(fromLeft, graph[i][j+1]) {
    graph[i][j] = 'F'
  } else {
    panic("cannot replace s")
  }
}