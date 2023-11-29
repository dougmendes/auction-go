package auction

import (
	"fmt"
	"sync"
)

type Auction struct {
	Items  []Item
	mu     sync.Mutex
	nextID int
}

type Item struct {
	ID    int
	Name  string
	Value int
	Bids  []Bid
}
type Bid struct {
	AuctionID int
	BidOwner  string
	Value     int
}

func NewAuction() *Auction {
	return &Auction{}
}

func (a *Auction) AddItem(name string, value int) {
	item := Item{
		ID:    a.nextID,
		Name:  name,
		Value: value,
	}
	a.nextID++
	a.Items = append(a.Items, item)
}

func (a *Auction) ListItems() {
	fmt.Println("Items in the auction:")
	for _, item := range a.Items {
		fmt.Printf("ID: %d, Name: %s, Value: %d\n", item.ID, item.Name, item.Value)
	}
}

func (item *Item) AddBid(auctionID int, bidOwner string, value int) {
	newBid := Bid{
		AuctionID: auctionID,
		BidOwner:  bidOwner,
		Value:     value,
	}
	item.Bids = append(item.Bids, newBid)
}
