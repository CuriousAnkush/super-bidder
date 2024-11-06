package domain

type Seller interface {
	Auctions(minAmount int, maxAmount int, actionName string) *Auction
	UniqueId() string
}

type DefaultSeller struct {
	sellerId string
}

func NewSeller(sellerId string) Seller {
	return &DefaultSeller{sellerId: sellerId}
}

func (receiver *DefaultSeller) Auctions(minAmount int, maxAmount int, auctionName string) *Auction {
	return NewAuction(receiver, auctionName, minAmount, maxAmount)
}

func (receiver *DefaultSeller) UniqueId() string {
	return receiver.sellerId
}
