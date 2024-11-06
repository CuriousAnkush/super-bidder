package infra

import "super-bidder/domain"

type BuyerRepository struct {
	db map[string]domain.Buyer
}

func NewBuyerRepository() domain.BuyerRepository {
	return &BuyerRepository{db: make(map[string]domain.Buyer)}
}

func (b *BuyerRepository) Create(buyer domain.Buyer) error {
	b.db[buyer.UniqueId()] = buyer

	return nil
}

func (b *BuyerRepository) Get(buyerId string) (domain.Buyer, error) {
	return b.db[buyerId], nil
}

func (b *BuyerRepository) Update(buyer domain.Buyer) (domain.Buyer, error) {
	b.db[buyer.UniqueId()] = buyer

	return buyer, nil
}
