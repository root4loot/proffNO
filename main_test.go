package proffno

import (
	"testing"
)

func TestFetchSubsidiaries(t *testing.T) {
	orgName := "DnB Bank ASA"

	results, err := GetSubsidiaries(orgName)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if results == nil {
		t.Fatalf("Expected results, got nil")
	}

	expectedSubsidiary := "DnB Asset Management AS"
	found := findSubsidiary(results.Tree, expectedSubsidiary)

	if !found {
		t.Errorf("Expected to find %s in subsidiaries, but it was not found", expectedSubsidiary)
	}
}

func findSubsidiary(tree Subsidiary, name string) bool {
	if tree.Name == name {
		return true
	}
	for _, sub := range tree.Sub {
		if findSubsidiary(sub, name) {
			return true
		}
	}
	return false
}
