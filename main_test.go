package main

import (
	"testing"
)

func TestSha256sum(t *testing.T) {
	// Test for existing file
	testFile := "README.md"
	expectedHash := "01b8b3c119c050e78a581b7d0918ca1e6f75f59af263e90e595f08a64543987b"
	hash, err := sha256sum(testFile)
	if err != nil {
		t.Errorf("Error calculating hash: %v", err)
	}
	if hash != expectedHash {
		t.Errorf("Expected %s, got %s", expectedHash, hash)
	}

	// Test for non-existent file
	nonExistentFile := "non_existent_file.txt"
	hash, err = sha256sum(nonExistentFile)
	if err == nil {
		t.Errorf("Expected error for non-existent file, got nil")
	}
	if hash != "" {
		t.Errorf("Expected empty hash for non-existent file, got %s", hash)
	}
}
