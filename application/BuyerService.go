package application

import (
	"super-bidder/domain"
)

type BuyerService struct {
	domain.BuyerRepository
}

func (receiver BuyerService) AddBuyer(buyerName string, budget int) error {
	bu := domain.NewBuyer(buyerName, budget)
	repo := receiver.BuyerRepository
	if err := repo.Create(bu); err != nil {
		return err
	}

	return nil
}

func (receiver BuyerService) UpdateBudget(buyerName string) error {
	repo := receiver.BuyerRepository
	if bu, err := repo.Get(buyerName); err != nil {
		if _, err := repo.Update(bu); err != nil {
			return err
		}
	}

	return nil
}
