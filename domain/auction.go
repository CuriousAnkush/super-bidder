package domain

type Auction struct {
	bids map[string]*Bid
	Seller
	name   string
	MinBid int
	MaxBid int
}

func NewAuction(seller Seller, name string, min int, max int) *Auction {
	return &Auction{
		Seller: seller,
		name:   name,
		MinBid: min,
		MaxBid: max,
		bids:   make(map[string]*Bid),
	}
}

func (a *Auction) UniqueId() string {
	return a.name
}

func (a *Auction) AddBid(b *Bid) {
	a.bids[b.UniqueId()] = b
}

func (a *Auction) WithdrawBid(b *Bid) {
	delete(a.bids, b.UniqueId())
}

func (a *Auction) Bids() map[string]*Bid {
	return a.bids
}

func (a *Auction) Close(strategy BidWinnerStrategy) *Bid {
	var bids []*Bid
	for _, bid := range a.bids {
		bids = append(bids, bid)
	}

	return strategy.FindWinner(bids)
}
