package application

import (
	"super-bidder/domain"
	"super-bidder/infra"
)

type SellerService struct {
	domain.SellerRepository
}

func NewSellerService() *SellerService {
	return &SellerService{SellerRepository: infra.NewSellerRepository()}
}

func (receiver *SellerService) AddSeller(sellerName string) error {
	bu := domain.NewSeller(sellerName)
	if err := receiver.SellerRepository.Create(bu); err != nil {
		return err
	}

	return nil
}
