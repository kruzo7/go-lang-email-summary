package formatter

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	f := New()
	if len(f.emailsnumber) != 0 {
		t.Errorf("Expected empty emailsnumber map, got %v", f.emailsnumber)
	}
}

// TestAdd tests the Add method of the Formatter struct
func TestAdd(t *testing.T) {
	f := New()
	f.Add("example.com")
	if f.emailsnumber["example.com"] != 1 {
		t.Errorf("Expected count 1 for example.com, got %d", f.emailsnumber["example.com"])
	}

	f.Add("example.com")
	if f.emailsnumber["example.com"] != 2 {
		t.Errorf("Expected count 2 for example.com, got %d", f.emailsnumber["example.com"])
	}

	f.Add("example2.com")
	if f.emailsnumber["example2.com"] != 1 {
		t.Errorf("Expected count 1 for example2.com, got %d", f.emailsnumber["example2.com"])
	}
}

// TestSortEmailsByValueDesc tests the sorting of emails by value in descending order
func TestSortEmailsByValueDesc(t *testing.T) {
	f := New()
	f.Add("example.com")
	f.Add("example.com")
	f.Add("example2.com")

	sorted := f.SortEmailsByValueDesc()

	if len(sorted) != 2 {
		t.Errorf("Expected 3 sorted pairs, got %d", len(sorted))
	}

	if sorted[0].Key != "example.com" || sorted[0].Value != 2 {
		t.Errorf("Expected example.com with count 2 at index 0, got %s with count %d", sorted[0].Key, sorted[0].Value)
	}
	if sorted[1].Key != "example2.com" || sorted[1].Value != 1 {
		t.Errorf("Expected example2.com with count 1 at index 1, got %s with count %d", sorted[1].Key, sorted[1].Value)
	}
}

// TestOutputConsole tests the output to console
func TestOutputConsole(t *testing.T) {
	f := New()
	f.Add("example.com")
	f.Add("example.com")
	f.Add("example2.com")

	var buf bytes.Buffer
	// Redirecting output to buffer
	r, w, _ := os.Pipe()
	old := os.Stdout
	defer func() { os.Stdout = old }()
	os.Stdout = w

	f.Print("console")

	w.Close()
	io.Copy(&buf, r)

	output := buf.String()
	expectedOutput := "example.com 2\nexample2.com 1\n"
	if output != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output)
	}
}

// TestOutputFile tests the output to a file
func TestOutputFile(t *testing.T) {
	f := New()
	f.Add("example.com")
	f.Add("example.com")
	f.Add("example2.com")

	// Call Print with "file" to create a file
	filename := f.Print("file")

	// Check if the file was created
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("Expected output file to be created, but it does not exist")
	}

	if err := os.Remove(filename); err != nil {
		t.Errorf("Failed to remove file %s: %v", filename, err)
	}
}
