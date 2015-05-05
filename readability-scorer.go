// readability-scorer scores the readability of text according to various formulas.
package main

import (
	"fmt"
	"github.com/BluntSporks/calculation"
	"github.com/BluntSporks/readability"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func main() {
	// Check arguments.
	if len(os.Args) == 1 {
		log.Fatal("Missing filename argument")
	}

	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)

	// Score ARI.
	ariScore := read.Ari(text)
	fmt.Printf("Automated Readability: %0.2f\n", ariScore)

	// Score CLI.
	cliScore := read.Cli(text)
	fmt.Printf("Coleman-Liau: %0.2f\n", cliScore)

	// Score Flesch-Kincaid.
	fkScore := read.Fk(text)
	fmt.Printf("Flesch-Kincaid: %0.2f\n", fkScore)

	// Score Gunning fog.
	gfiScore := read.Gfi(text)
	fmt.Printf("Gunning fog: %0.2f\n", gfiScore)

	// Score SMOG.
	smogScore := read.Smog(text)
	fmt.Printf("SMOG: %0.2f\n", smogScore)

	// Yield average score.
	values := []float64{ariScore, cliScore, fkScore, gfiScore, smogScore}
	sort.Float64s(values)
	avg := calc.Mean(values)
	stddev := calc.StdDev(values)
	fmt.Printf("Sorted scores: [%0.2f, %0.2f, %0.2f, %0.2f, %0.2f]\n", values[0], values[1], values[2], values[3], values[4])
	fmt.Printf("Average score: %0.2f\n", avg)
	fmt.Printf("Std Dev of scores: %0.2f\n", stddev)
}
