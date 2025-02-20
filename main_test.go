package proffno

import (
	"testing"
)

func TestFetchSubsidiaries(t *testing.T) {
	orgName := "DnB Bank ASA"
	depth := 2
	minOwnership := 50.0

	corpData, err := FetchSubsidiaries(orgName, depth, minOwnership)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if corpData == nil {
		t.Fatalf("Expected corporate data, got nil")
	}

	expectedSubsidiary := "DnB Asset Management AS"
	found := findSubsidiary(corpData.Tree, expectedSubsidiary)

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
