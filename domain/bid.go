package domain

import "fmt"

type Bid struct {
	Buyer     Buyer
	auction   *Auction
	BidAmount int
}

func NewBid(buyer Buyer, auction *Auction, bidAmount int) (*Bid, error) {
	if buyer.Balance() < bidAmount {
		return nil, fmt.Errorf("buyer does not have enough balance")
	}

	if auction.MinBid > bidAmount {
		return nil, fmt.Errorf("bid amount is less than minimum amount")
	}

	if auction.MaxBid < bidAmount {
		return nil, fmt.Errorf("bid amount is more than maximum amount")
	}

	return &Bid{
		Buyer:     buyer,
		auction:   auction,
		BidAmount: bidAmount,
	}, nil
}

func (receiver Bid) UniqueId() string {
	auction := receiver.auction
	return auction.name + receiver.Buyer.UniqueId()
}
