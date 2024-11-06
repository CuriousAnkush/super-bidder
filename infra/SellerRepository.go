package infra

import "super-bidder/domain"

type SellerRepository struct {
	db map[string]domain.Seller
}

func NewSellerRepository() domain.SellerRepository {
	return &SellerRepository{db: make(map[string]domain.Seller)}
}

func (b SellerRepository) Create(seller domain.Seller) error {
	bu := seller
	b.db[bu.UniqueId()] = seller

	return nil
}

func (b SellerRepository) Get(sellerId string) (domain.Seller, error) {
	return b.db[sellerId], nil
}
