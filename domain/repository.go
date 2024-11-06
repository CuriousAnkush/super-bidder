package domain

type AuctionRepository interface {
	Create(auction *Auction) error
	Get(auctionName string) (*Auction, error)
	GetBids(auctionID string) []*Bid
	Update(auction *Auction) error
}

type BuyerRepository interface {
	Create(buyer Buyer) error
	Get(buyerId string) (Buyer, error)
	Update(buyer Buyer) (Buyer, error)
}

type SellerRepository interface {
	Create(seller Seller) error
	Get(sellerId string) (Seller, error)
}

type BidRepository interface {
	Create(bid *Bid) error
	Update(bid *Bid) error
	Delete(bid *Bid) error
	Get(bid string) (*Bid, error)
}
