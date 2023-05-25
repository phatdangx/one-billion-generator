package main

import (
	sol "one-billion-generator/solutions"
)

func main() {
	const (
		oneHundredThousand = 100000
		oneMillion         = 1000000
		oneHundredMillions = 100000000
		oneBillion         = 1000000000
	)

	batchSize := oneHundredMillions
	goalNumber := oneBillion

	// Execute solution 1
	s1 := sol.NewSolution1(goalNumber, batchSize)
	s1.Execute()
}
