package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var dirs = [][]int {
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1},  /*   */  {0, 1},
	{1, -1},  {1, 0},  {1, 1},
}

func main() {
	f := "./3/b/input"
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

	numsPerGear := make(map[string][]int)
	for i := range(rows) {
		for j := 0; j<len(rows[i]); j++ {
			curNum := 0
			gears := make(map[string]bool)
			for j<len(rows[i]) && unicode.IsDigit(rows[i][j]) {
				curNum = 10*curNum + int(rows[i][j] - '0')
				addBorderingSymbol(rows, i, j, gears)
				j++
			}
			for g := range(gears) {
				numsPerGear[g] = append(numsPerGear[g], curNum)
			}
		}
	}

	// fmt.Println("GEARS: ", numsPerGear)
	res := 0
	for _, els := range(numsPerGear) {
		if len(els) != 2 {
			continue
		}
		res += els[0] * els[1]
	}
	fmt.Println("res ", res)
}

func addBorderingSymbol(rows [][]rune, i, j int, m map[string]bool) {
	for _, dir := range(dirs) {
		newi := i + dir[0]
		newj := j + dir[1]
		if newi < 0 || newi == len(rows) || newj < 0 || newj == len(rows[0]) {
			continue
		}
		if isSymbol(rows[newi][newj]) {
			s := strconv.Itoa(newi) + "," + strconv.Itoa(newj)
			m[s] = true
		}
	}
}

func isSymbol(r rune) bool {
	return !(unicode.IsDigit(r) || r == '.')
}