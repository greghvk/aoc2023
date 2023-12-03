package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var dirs = [][]int {
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1},  /*   */  {0, 1},
	{1, -1},  {1, 0},  {1, 1},
}

func main() {
	f := "./3/a/input"
	file, err := os.Open(f)
	if err != nil {
			fmt.Println(err)
			return
	}
	defer file.Close()

	rows := make([][]rune, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, []rune(scanner.Text()))
	}

	res := 0
	for i := range(rows) {
		for j := 0; j<len(rows[i]); j++ {
			curNum := 0
			ok := false
			for j<len(rows[i]) && unicode.IsDigit(rows[i][j]) {
				curNum = 10*curNum + int(rows[i][j] - '0')
				if !ok && bordersSymbol(rows, i, j) { // !ok to skip unnecessary processing
					ok = true
				}
				j++
			}
			if ok {
				res += curNum
				fmt.Println("counting ", curNum)
			} else if curNum != 0 {
				fmt.Println("skipping ", curNum)
			}
		}
	}
	fmt.Println("res ", res)
}

func bordersSymbol(rows [][]rune, i, j int) bool {
	for _, dir := range(dirs) {
		newi := i + dir[0]
		newj := j + dir[1]
		if newi < 0 || newi == len(rows) || newj < 0 || newj == len(rows[0]) {
			continue
		}
		if isSymbol(rows[newi][newj]) {
			return true
		}
	}
	return false
}

func isSymbol(r rune) bool {
	return !(unicode.IsDigit(r) || r == '.')
}