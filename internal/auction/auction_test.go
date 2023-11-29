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

	expectedOutput := "Items in the auction:\nID: 0, Name: Item 1, Value: 100\nID: 1, Name: Item 2, Value: 200\n"

	output := captureOutput(func() {
		a.ListItems()
	})

	if output != expectedOutput {
		t.Errorf("Expected output:\n%s\nBut got:\n%s", expectedOutput, output)
	}
}

func TestAddBid(t *testing.T) {
	item := Item{}
	item.AddBid(1, "User1", 50)

	if len(item.Bids) != 1 {
		t.Errorf("Expected 1 bid for the item, but got %d", len(item.Bids))
	}

	if item.Bids[0].AuctionID != 1 {
		t.Errorf("Expected AuctionID 1, but got %d", item.Bids[0].AuctionID)
	}

	if item.Bids[0].BidOwner != "User1" {
		t.Errorf("Expected BidOwner 'User1', but got '%s'", item.Bids[0].BidOwner)
	}

	if item.Bids[0].Value != 50 {
		t.Errorf("Expected Value 50, but got %d", item.Bids[0].Value)
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
