package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/root4loot/proffno"
)

func main() {
	results, err := proffno.GetSubsidiaries("DnB Bank ASA")
	if err != nil {
		log.Fatal(err)
	}

	if results == nil {
		log.Println("No subsidiaries found")
		return
	}

	var processSubsidiaries func(sub proffno.Subsidiary)

	processSubsidiaries = func(sub proffno.Subsidiary) {
		indent := strings.Repeat("  ", sub.Depth)
		fmt.Printf("%s%d. %s (%.2f%%)\n", indent, sub.Depth+1, sub.Name, sub.OwnedPercentage)

		for _, child := range sub.Sub {
			processSubsidiaries(child)
		}
	}

	processSubsidiaries(results.Tree)
}
