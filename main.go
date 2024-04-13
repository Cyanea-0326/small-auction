package main

import (
	"fmt"
	"time"
	"io"
	"net/http"
	"test-mutex/logic"

	"github.com/gorilla/mux"
)

func makeAuction(id string) (*logic.Auction) {
	auction := &logic.Auction{
		ID:			id,
		StartTime:	time.Now(),
	}
	return auction
}

func makeBidder(id string, deposit float64) (*logic.Bidder) {
	bidder := &logic.Bidder{
		ID:			id,
		Deposit:	10000,
		Bids:		make(map[string]float64),
	}
	return bidder
}

func PlaceBidInAuction(a *logic.Auction, b *logic.Bidder, amount float64) error {	
	err := b.PlaceBid(a, amount)
	if err != nil {
		fmt.Println("オークションでの入札エラー: ", err)
	}
	return nil
}

///////////////////////////////////////////////////////////

func auctionHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	auctionID := string(body)
	a1 := makeAuction(auctionID);
	a1.Instance(time.Second * 60);
}

func bidHandler(w http.ResponseWriter, r *http.Request) {
	// b1 := makeBidder("B1", 10000);
	// b2 := makeBidder("B2", 5000);

	// PlaceBidInAuction(a1, b1, 3000)
	// PlaceBidInAuction(a1, b2, 6000)
	// PlaceBidInAuction(a1, b1, 4000)
	// PlaceBidInAuction(a1, b2, 7000)
}


func main() {
	router := mux.NewRouter()
	fmt.Println("Server started at :8080");

	router.HandleFunc("/auction", auctionHandler).Methods("POST")

	http.ListenAndServe(":8080", router);
}