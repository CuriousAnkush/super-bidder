package infra

import "super-bidder/domain"

type BidRepository struct {
	db map[string]*domain.Bid
}

func NewBidRepository() domain.BidRepository {
	return &BidRepository{db: make(map[string]*domain.Bid)}
}

func (b BidRepository) Create(bid *domain.Bid) error {
	b.db[bid.UniqueId()] = bid
	return nil
}

func (b BidRepository) Update(bid *domain.Bid) error {
	b.db[bid.UniqueId()] = bid

	return nil
}

func (b BidRepository) Delete(bid *domain.Bid) error {
	delete(b.db, bid.UniqueId())

	return nil
}

func (b BidRepository) Get(bid string) (*domain.Bid, error) {
	return b.db[bid], nil
}
