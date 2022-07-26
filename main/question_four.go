package main

import "fmt"

const (
	concave = 0
	convex  = 1
	smooth  = 2
)

type puzzle struct {
	top   int
	below int
	left  int
	right int
}

func validPuzzles(puzzles [3][3]puzzle) bool {

	xlen := len(puzzles)
	ylen := len(puzzles[0])
	for i := 0; i < xlen; i++ {
		for j := 0; j < ylen; j++ {
			fmt.Println(i, j)
			item := puzzles[i][j]
			if i == 0 && item.top != smooth {
				return false
			}
			if j == 0 && item.left != smooth {
				return false
			}
			if j == ylen-1 && item.right != smooth {
				return false
			}
			if i == xlen-1 && item.below != smooth {
				return false
			}

			var rightPuzzle puzzle
			var belowPuzzle puzzle
			if j != ylen-1 {
				rightPuzzle = puzzles[i][j+1]
			}
			if i != xlen-1 {
				belowPuzzle = puzzles[i+1][j]
			}
			if (rightPuzzle != puzzle{}) {
				if item.right+rightPuzzle.left != 1 {
					return false
				}
			}
			if (belowPuzzle != puzzle{}) {
				if item.below+belowPuzzle.top != 1 {
					return false
				}
			}

		}
	}
	return true
}

func main() {
	puzzle1 := puzzle{top: smooth, left: smooth, right: concave, below: convex}
	puzzle2 := puzzle{top: smooth, left: convex, right: concave, below: convex}
	puzzle3 := puzzle{top: smooth, left: convex, right: smooth, below: convex}

	puzzle4 := puzzle{top: concave, left: smooth, right: concave, below: convex}
	puzzle5 := puzzle{top: concave, left: convex, right: concave, below: convex}
	puzzle6 := puzzle{top: concave, left: convex, right: smooth, below: convex}

	puzzle7 := puzzle{top: concave, left: smooth, right: concave, below: smooth}
	puzzle8 := puzzle{top: concave, left: convex, right: concave, below: smooth}
	puzzle9 := puzzle{top: concave, left: convex, right: smooth, below: smooth}

	puzzles := [3][3]puzzle{}
	puzzles[0][0] = puzzle1
	puzzles[0][1] = puzzle2
	puzzles[0][2] = puzzle3
	puzzles[1][0] = puzzle4
	puzzles[1][1] = puzzle5
	puzzles[1][2] = puzzle6
	puzzles[2][0] = puzzle7
	puzzles[2][1] = puzzle8
	puzzles[2][2] = puzzle9

	fmt.Println(validPuzzles(puzzles))

}
