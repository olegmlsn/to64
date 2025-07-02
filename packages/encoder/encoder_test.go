package encoder

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEncodeFileToBase64(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	content := "Hello, test!"

	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}

	got, err := EncodeFileToBase64(testFile)
	if err != nil {
		t.Fatalf("error during encoding: %v", err)
	}

	want := "SGVsbG8sIHRlc3Qh"
	if got != want {
		t.Errorf("incorrect result:\ngot: %s\nexpected: %s", got, want)
	}
}

func TestFileNotFound(t *testing.T) {
	_, err := EncodeFileToBase64("nonexistent.file")
	if err == nil {
		t.Error("expected an error for missing file")
	}
}
