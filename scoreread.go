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
	useFk := flag.Bool("fk", false, "Use Flesch-Kincaid Grade Level")
	flag.Parse()

	// Check arguments.
	if *file == "" {
		log.Fatal("Missing -file argument")
	}
	if !*useAri || !*useFk {
		// Assume all tests are requested if no tests are requested.
		*useAri = true
		*useFk = true
	}

	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)

	// Score ARI, if requested.
	if *useAri {
		ariScore := read.Ari(text)
		fmt.Printf("ARI: %0.2f\n", ariScore)
	}

	// Score Flesch-Kincaid, if requested.
	if *useFk {
		fkScore := read.Fk(text)
		fmt.Printf("Flesch-Kincaid: %0.2f\n", fkScore)
	}
}
