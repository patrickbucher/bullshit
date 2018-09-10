package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sync"

	"github.com/patrickbucher/bullshit"
)

const (
	outputDir = "output"
)

type Bullshit struct {
	Rows []Row
}

type Row struct {
	Cols []Col
}

type Col struct {
	Term string
}

var cards = flag.Int("n", 1, "Number of Bullshit Bingo cards to be generated")
var cols = flag.Int("c", 4, "Number of columns for each Bullshit Bingo card")
var rows = flag.Int("r", 4, "Number of rows for each Bullshit Bingo card")

var wg sync.WaitGroup

func main() {
	flag.Parse()
	if *rows < 1 || *cols < 1 || *cards < 1 {
		fatal("parameters r, c and n must be positive\n")
	}
	amount := *rows * *cols
	s := bufio.NewScanner(os.Stdin)
	var terms []string
	for s.Scan() {
		terms = append(terms, s.Text())
	}
	if len(terms) < amount {
		fatal("%d terms or more needed\n", amount)
	}
	if err := os.RemoveAll(outputDir); err != nil {
		fatal("removing %s: %v\n", outputDir, err)
	}
	if err := os.Mkdir(outputDir, 0755); err != nil {
		fatal("creating %s: %v\n", outputDir, err)
	}
	for i := 0; i < *cards; i++ {
		go produceBullshit(terms, i)
		wg.Add(1)
	}
	wg.Wait()
}

func produceBullshit(terms []string, fileNumber int) {
	w := int(math.Log10(float64(*cards))) + 1
	format := "bullshit%0*d.html"
	filename := outputDir + string(os.PathSeparator) +
		fmt.Sprintf(format, w, fileNumber+1)
	file, err := os.Create(filename)
	if err != nil {
		fatal("creating file %s: %v\n", filename, err)
	}
	defer file.Close()
	bs := arrange(terms, *rows, *cols)
	bullshit.BullshitTemplate.Execute(file, bs)
	wg.Done()
}

func arrange(terms []string, rows, cols int) Bullshit {
	items := bullshit.Shuffle(terms)
	var bullshit Bullshit
	bsRows := make([]Row, rows)
	for r := 0; r < rows; r++ {
		bsCols := make([]Col, cols)
		for c := 0; c < cols; c++ {
			bsCols[c] = Col{items[r*rows+c]}
		}
		bsRows[r] = Row{bsCols}
	}
	bullshit.Rows = bsRows
	return bullshit
}

func fatal(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
