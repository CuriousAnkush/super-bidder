package infra

import "super-bidder/domain"

type AuctionRepository struct {
	db map[string]*domain.Auction
}

func NewAuctionRepository() domain.AuctionRepository {
	return &AuctionRepository{db: make(map[string]*domain.Auction)}
}

func (a *AuctionRepository) Create(auction *domain.Auction) error {
	a.db[auction.UniqueId()] = auction

	return nil
}

func (a *AuctionRepository) Update(auction *domain.Auction) error {
	a.db[auction.UniqueId()] = auction

	return nil
}

func (a *AuctionRepository) Get(auctionId string) (*domain.Auction, error) {
	return a.db[auctionId], nil
}

func (a *AuctionRepository) GetBids(auctionID string) []*domain.Bid {
	//return a.db[auctionID].Bids()

	return []*domain.Bid{}
}
