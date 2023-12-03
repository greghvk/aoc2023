package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f := "./2/a/input"
	file, err := os.Open(f)
	if err != nil {
			fmt.Println(err)
			return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res := 0
	i := 1
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), ":") 
		if analyzeParts(strings.Split(str[1], ";")) {
			res += i
		}
		i++
	}
	fmt.Println("res ", res)
}

func analyzeParts(parts []string) bool {
	for _, part := range(parts) {
		nums := strings.Split(part, ",")
		if !possible(nums) {
			return false
		}
	}
	return true
}

var m = map[string]int{
	"red": 12,
	"green": 13,
	"blue": 14,
}

func possible(nums []string) bool {
	for _, els := range(nums) {
		el := strings.Split(els, " ")
		color := el[2]
		val, err := strconv.Atoi(el[1])
		if err != nil {
			fmt.Println(err)
			return false
		}
		if m[color] < val {
			return false
		}
	}
	return true
}