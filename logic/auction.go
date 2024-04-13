package logic

import (
	"fmt"
	"time"
	"errors"
)

func (bidder *Bidder) GetTotalBidAmount() float64 {
	bidder.Mutex.Lock();
	defer bidder.Mutex.Unlock();

	totalBidAmount := 0.0
	for _, amount := range bidder.Bids {
		totalBidAmount += amount
	}
	return totalBidAmount
}

func (a *Auction) Instance(duration time.Duration) {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()

	a.StartTime = time.Now()
	a.EndTime = a.StartTime.Add(duration)
	a.Active = true
	fmt.Printf("------------------------------------------------\n\n")
	fmt.Printf("Auction %s started\n", a.ID,)
	fmt.Printf("will end at %s\n\n", a.EndTime)
	fmt.Printf("------------------------------------------------\n")

	go func() {
		time.Sleep(duration)
		a.CloseAuction()
	}()
}

// オークションの入札プロセス
func (a *Auction) PlaceBid(bid Bid) error {
	// オークションの終了時間内に入札できるようにする
	if time.Now().After(a.EndTime) {
		return fmt.Errorf("オークションは終了しました")
	}

	a.Mutex.Lock()
	defer a.Mutex.Unlock()

	// 入札額を追加
	// 最高額入札を更新
	a.Bids = append(a.Bids, bid)
	if bid.Amount > a.HighestBid.Amount {
		a.HighestBid = bid
	}
	return nil
}

func (a *Auction) IsActive() bool {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	return a.Active
}

// オークションの終了プロセス
func (a *Auction) CloseAuction() {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()

	if a.Active == false {
		return ;
	}
	a.Active = false
	fmt.Printf("オークション %s は終了しました\n", a.ID)
	fmt.Printf("最高額入札: 入札者 %s、金額 %.2f\n", a.HighestBid.BidderID, a.HighestBid.Amount)
}

// 入札者がオークションに入札するプロセス
func (b *Bidder) PlaceBid(auction *Auction, amount float64) error {
	if !auction.Active {
		return errors.New("オークションは終了しています")
	}
	if b.Deposit < amount {
		return fmt.Errorf("デポジットが不足しています")
	}

	// 入札プロセス
	// デポジットを減らす
	// 入札情報を記録
	err := auction.PlaceBid(Bid{BidderID: b.ID, Amount: amount})
	if err != nil {
		return err
	}
	b.Deposit -= amount
	b.Bids[auction.ID] = amount
	return nil
}


