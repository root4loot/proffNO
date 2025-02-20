package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/root4loot/proffno"
)

func main() {
	// Fetch subsidiaries of a company and its sub-subsidiaries owned by more than 50%
	corpData, err := proffno.FetchSubsidiaries("DnB Bank ASA", 2, 50)
	if err != nil {
		log.Fatalf("Failed to fetch subsidiaries: %v", err)
	}

	if corpData == nil {
		log.Println("No subsidiaries found")
		return
	}

	printSubsidiaries(corpData.Tree, 0)
}

// Recursively print subsidiaries
func printSubsidiaries(sub proffno.Subsidiary, level int) {
	fmt.Printf("%s%s (%.2f%%)\n", strings.Repeat("  ", level), sub.Name, sub.OwnedPercentage)
	for _, child := range sub.Sub {
		printSubsidiaries(child, level+1)
	}
}
