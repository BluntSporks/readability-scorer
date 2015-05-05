package main

import (
	"flag"
	"fmt"
	"github.com/BluntSporks/readability"
	"io/ioutil"
	"log"
)

func main() {
	file := flag.String("file", "", "Name of file to filter")
	useAri := flag.Bool("ari", false, "Use Automated Readbility Index (ARI)")
	useCli := flag.Bool("cli", false, "Use Coleman-Liau Index (CLI)")
	useFk := flag.Bool("fk", false, "Use Flesch-Kincaid Grade Level")
	useGfi := flag.Bool("gfi", false, "Use Gunning fog index")
	flag.Parse()

	// Check arguments.
	if *file == "" {
		log.Fatal("Missing -file argument")
	}
	if !*useAri && !*useCli && !*useFk && !*useGfi {
		// Assume all tests are requested if no tests are requested.
		*useAri = true
		*useCli = true
		*useFk = true
		*useGfi = true
	}

	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)

	// Score ARI, if requested.
	if *useAri {
		ariScore := read.Ari(text)
		fmt.Printf("Automated Readability: %0.2f\n", ariScore)
	}

	// Score CLI, if requested.
	if *useCli {
		cliScore := read.Cli(text)
		fmt.Printf("Coleman-Liau: %0.2f\n", cliScore)
	}

	// Score Flesch-Kincaid, if requested.
	if *useFk {
		fkScore := read.Fk(text)
		fmt.Printf("Flesch-Kincaid: %0.2f\n", fkScore)
	}

	// Score Gunning fog, if requested.
	if *useGfi {
		gfiScore := read.Gfi(text)
		fmt.Printf("Gunning fog: %0.2f\n", gfiScore)
	}
}
