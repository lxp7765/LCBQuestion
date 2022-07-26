package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rand13() int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(13) + 1
	return r
}

func rand5() int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	return r
}

func rand13ToRand5() int {

	for {
		r65 := (rand13()-1)*5 + rand13()
		if r65 <= 165 {
			result := r65%5 + 1
			return result
		}
		r65 = (r65-165-1)*5 + rand13()
		if r65 <= 50 {
			result := r65%5 + 1
			return result
		}
	}
}

func rand5ToRand13() int {
	for {
		r25 := (rand5()-1)*5 + rand5()
		if r25 <= 13 {
			result := r25%13 + 1
			return result
		}

	}

}

func main() {
	fmt.Println(rand13ToRand5())
	fmt.Println(rand5ToRand13())
}
