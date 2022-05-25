package main // import "adduntilmirror"

import (
	"fmt"
	"strconv"
)

func getMirrorNumber(n int) int {
	nStr := fmt.Sprint(n)

	newNstr := ""
	for _, s := range nStr {
		newNstr = string(s) + newNstr
	}

	newN, _ := strconv.Atoi(newNstr)

	return newN
}

func isMirror(n, m int) bool {
	result := true

	nStr := fmt.Sprint(n)
	mStr := fmt.Sprint(m)

	for i := range nStr {
		if nStr[i] != mStr[i] {
			result = false
			break
		}
	}

	return result
}

func add(a, b int) int {
	return a + b
}

func main() {
	out := 0
	num := 123
	var s1 string

	fmt.Print("Please enter a number:")
	_, _ = fmt.Scan(&s1)

	num, _ = strconv.Atoi(s1)

	for {
		mir := getMirrorNumber(num)
		
		if isMirror(num, mir) || out > 999999999 {
			break
		}

		out = add(num, mir)

		fmt.Println(out)

		num = out
	}
}
