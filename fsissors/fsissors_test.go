package fsissors_test

import (
	"github.com/skiloop/fsissors/fsissors"
	"os"
	"testing"
)

func TestFileCopy(t *testing.T) {
	source := "test_source.txt"
	target := "test_target.txt"
	content := []byte("This is a test file.")

	// Create a test source file
	err := os.WriteFile(source, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}
	defer os.Remove(source)
	defer os.Remove(target)

	// Perform file copy
	err = fsissors.FileCopy(source, 0, target, 0, 1024, int64(len(content)))
	if err != nil {
		t.Fatalf("FileCopy failed: %v", err)
	}

	// Verify the target file content
	targetContent, err := os.ReadFile(target)
	if err != nil {
		t.Fatalf("Failed to read target file: %v", err)
	}
	if string(targetContent) != string(content) {
		t.Errorf("File content mismatch. Expected: %s, Got: %s", content, targetContent)
	}
}
func TestFileTruncate(t *testing.T) {
	file := "test_truncate.txt"
	content := []byte("This is a test file.")

	// Create a test file
	err := os.WriteFile(file, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(file)

	// Perform file truncation
	err = fsissors.FileTruncate(file, 10)
	if err != nil {
		t.Fatalf("FileTruncate failed: %v", err)
	}

	// Verify the file size
	info, err := os.Stat(file)
	if err != nil {
		t.Fatalf("Failed to stat file: %v", err)
	}
	if info.Size() != 10 {
		t.Errorf("File size mismatch. Expected: 10, Got: %d", info.Size())
	}
}
func TestBytesModify(t *testing.T) {
	file := "test_modify.txt"
	content := []byte("This is a test file.")
	modifiedContent := []byte("This is a TEST file.")

	// Create a test file
	err := os.WriteFile(file, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(file)

	// Perform byte modification
	err = fsissors.BytesModify(file, 10, 4, 4, "54455354") // "TEST" in hex
	if err != nil {
		t.Fatalf("BytesModify failed: %v", err)
	}

	// Verify the file content
	result, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	if string(result) != string(modifiedContent) {
		t.Errorf("File content mismatch. Expected: %s, Got: %s", modifiedContent, result)
	}
}
