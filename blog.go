package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func blogReport() {
	args := flag.Args()
	if len(args) != 0 {
		var err error
		goalWPD, err = strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
	}

	days := max(daysSinceLastYear(), 1)

	expectedProgress := goalWPD * days
	totalWords := totalWordCount()
	observedProgress := totalWords - initWords

	observedRate := float64(observedProgress) / float64(days)
	fmt.Println("BLOGRESS REPORT:")

	if observedProgress < expectedProgress {
		fmt.Printf("  %d words behind schedule\n", expectedProgress-observedProgress)
	}

	fmt.Printf("  %d%% of goal\n", int(observedRate/float64(goalWPD)*100))

}
