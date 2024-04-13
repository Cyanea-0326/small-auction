package logic

import (
	"time"
	"sync"
)

type Bid struct {
	BidderID string
	Amount   float64
}

// オークションを表す構造体
type Auction struct {
	ID			string
	Bids		[]Bid
	StartTime	time.Time
	EndTime		time.Time
	HighestBid	Bid
	Mutex		sync.Mutex // 同期用
	Active		bool
}

// 入札者を表す構造体
type Bidder struct {
	ID		string
	Deposit	float64
	Bids	map[string]float64 // オークションIDと入札額
	Mutex	sync.Mutex
}
