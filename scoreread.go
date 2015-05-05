package main

import (
	"fmt"
	"github.com/BluntSporks/calculation"
	"github.com/BluntSporks/readability"
	"io/ioutil"
	"log"
	"os"
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
	avg := calc.Mean(values)
	fmt.Printf("Average score: %0.2f\n", avg)
}
