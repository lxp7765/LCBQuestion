package main

import (
	"fmt"
)

// Perm() 对 a 形成的每⼀排列调⽤ f().
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// 对索引 i 从 0 到 len(a) - 1，实现递归函数 perm().
var input = "ABC"
var inputLength = len(input)

func perm(a []rune, f func([]rune), i int) {

	if i == inputLength {
		f(a[len(a)-inputLength:])
		return
	}
	dict := map[rune]string{}
	for n := inputLength; n < len(a); n++ {
		dict[a[n]] = ""
	}
	for n := 0; n < inputLength; n++ {
		if _, ok := dict[a[n]]; !ok {
			perm(append(a, a[n]), f, i+1)
		}
	}
}
func main() {
	Perm([]rune(input), func(a []rune) {
		fmt.Println(string(a))
	})
}
