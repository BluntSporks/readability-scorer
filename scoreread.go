package main

import (
	"flag"
	"fmt"
	"github.com/BluntSporks/calculation"
	"github.com/BluntSporks/readability"
	"io/ioutil"
	"log"
)

func main() {
	file := flag.String("file", "", "Name of file to filter")
	flag.Parse()

	// Check arguments.
	if *file == "" {
		log.Fatal("Missing -file argument")
	}

	bytes, err := ioutil.ReadFile(*file)
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
