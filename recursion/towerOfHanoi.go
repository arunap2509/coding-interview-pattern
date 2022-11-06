package recursion

import "fmt"

const NUM_OF_PEGS = 3
var result [][]int
var pegs [][]int

func TowerOfHanoi(numOfRings int) {

	pegs = append(pegs, generateSequence(numOfRings))

	for i := 1; i < NUM_OF_PEGS; i++ {
		pegs = append(pegs, []int{})
	}

	fmt.Println("pegs before computing", pegs)

	computeTowerOfHanoi(numOfRings, 0, 1, 2)

	fmt.Println("pegs after computing", pegs)
}

func generateSequence(num int) []int {
	var seq []int

	for i := num; i > 0; i-- {
		seq = append(seq, i)
	}

	return seq
}

func computeTowerOfHanoi(numRingsToMove, fromPeg, toPeg, usePeg int) {
	if numRingsToMove == 1 {
		pegs[toPeg] = append(pegs[toPeg], pop(&pegs[fromPeg]))
		fmt.Printf("moved ring from peg %d to peg %d \n", fromPeg, toPeg)
	} else {
		computeTowerOfHanoi(numRingsToMove - 1, fromPeg, usePeg, toPeg)
		pegs[toPeg] = append(pegs[toPeg], pop(&pegs[fromPeg]))
		fmt.Printf("moved ring from peg %d to peg %d \n", fromPeg, toPeg)
		computeTowerOfHanoi(numRingsToMove - 1, usePeg, toPeg, fromPeg)
	}
}

func pop(arr *[]int) int {
	len := len(*arr)

	val := (*arr)[len - 1]

	*arr = (*arr)[:len-1]

	return val
}