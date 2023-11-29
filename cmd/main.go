package main

import (
	"github.com/dougmendes/auction-go/internal/auction"
)

func main() {

	auction := auction.NewAuction()

	auction.AddItem("Watch", 100)
	auction.AddItem("Painting", 200)
	auction.AddItem("Laptop", 500)

	auction.ListItems()
}
