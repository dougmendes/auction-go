package auction

import "fmt"

type Item struct {
	Name  string
	Value int
}

type Auction struct {
	Items []Item
}

func NewAuction() *Auction {
	return &Auction{}
}

func (a *Auction) AddItem(name string, value int) {
	item := Item{Name: name, Value: value}
	a.Items = append(a.Items, item)
}

func (a *Auction) ListItems() {
	fmt.Println("Items in the auction:")
	for _, item := range a.Items {
		fmt.Printf("Name: %s, Value: %d\n", item.Name, item.Value)
	}
}
