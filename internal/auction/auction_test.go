package auction

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestAddItem(t *testing.T) {
	a := NewAuction()
	a.AddItem("Item 1", 100)

	if len(a.Items) != 1 {
		t.Errorf("Expected 1 item in the auction, but got %d", len(a.Items))
	}
}

func TestListItems(t *testing.T) {
	a := NewAuction()
	a.AddItem("Item 1", 100)
	a.AddItem("Item 2", 200)

	expectedOutput := "Items in the auction:\nName: Item 1, Value: 100\nName: Item 2, Value: 200\n"

	if testing.Verbose() {
		a.ListItems()
	}

	output := captureOutput(func() {
		a.ListItems()
	})

	if output != expectedOutput {
		t.Errorf("Expected output:\n%s\nBut got:\n%s", expectedOutput, output)
	}
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
