package solutions

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type Solution1 struct {
	goalNumber int
	batchSize  int
}

func NewSolution1(goalNumber, batchSize int) Solution1 {
	return Solution1{
		goalNumber: goalNumber,
		batchSize:  batchSize,
	}
}

func (s1 *Solution1) Execute() {
	// create file
	f, err := os.Create("res.txt")
	checkError(err)
	defer f.Close()

	startTime := time.Now()

	wg := sync.WaitGroup{}

	for i := 1; i <= s1.goalNumber; i += s1.batchSize {
		wg.Add(1)
		go writeToFile(i, i+s1.batchSize-1, &wg, f)
	}

	wg.Wait()

	// calculate running time
	duration := time.Since(startTime)
	fmt.Printf("%s\n", fmt.Sprintf("\n---\nSolution 1 running time: %s ", duration))
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func writeToFile(start, end int, wg *sync.WaitGroup, f *os.File) {
	defer wg.Done()
	fmt.Printf("Generate from %v to %v\n", start, end)
	w := bufio.NewWriter(f)
	for i := start; i <= end; i++ {
		w.WriteString(fmt.Sprintf("%v\n", i))
	}
	w.Flush()
	fmt.Printf("Finish generate from %v to %v\n", start, end)
}
