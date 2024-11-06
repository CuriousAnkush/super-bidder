package application

import (
	"super-bidder/domain"
)

type BidService struct {
	domain.BidRepository
	domain.AuctionRepository
	domain.BuyerRepository
}

func NewBidService(biR domain.BidRepository, aR domain.AuctionRepository, bR domain.BuyerRepository) *BidService {
	return &BidService{
		BidRepository:     biR,
		AuctionRepository: aR,
		BuyerRepository:   bR,
	}
}

func (receiver *BidService) CreateBid(buyerName string, auctionId string, amount int) error {
	buyer, err := receiver.BuyerRepository.Get(buyerName)
	if err != nil {
		return err
	}

	auction, err := receiver.AuctionRepository.Get(auctionId)
	if err != nil {
		return err
	}

	bid, err := domain.NewBid(buyer, auction, amount)
	if err != nil {
		return err
	}
	auction.AddBid(bid)
	aerr := receiver.AuctionRepository.Create(auction)
	if aerr != nil {
		return aerr
	}
	if err := receiver.BidRepository.Create(bid); err != nil {
		return err
	}

	return nil
}

func (receiver *BidService) UpdateBid(buyerName string, auctionId string, amount int) error {
	buyer, err := receiver.BuyerRepository.Get(buyerName)
	if err != nil {
		return err
	}

	auction, err := receiver.AuctionRepository.Get(auctionId)
	if err != nil {
		return err
	}

	newBid, errb := domain.NewBid(buyer, auction, amount)
	if errb != nil {
		return err
	}
	auction.AddBid(newBid)
	aerr := receiver.AuctionRepository.Update(auction)
	if aerr != nil {
		return aerr
	}
	if err := receiver.BidRepository.Update(newBid); err != nil {
		return err
	}

	return nil
}

func (receiver *BidService) Withdraw(buyerName string, auctionId string) error {
	buyer, err := receiver.BuyerRepository.Get(buyerName)
	if err != nil {
		return err
	}

	auction, err := receiver.AuctionRepository.Get(auctionId)
	if err != nil {
		return err
	}

	newBid, errb := domain.NewBid(buyer, auction, auction.MinBid)
	if errb != nil {
		return err
	}
	auction.WithdrawBid(newBid)
	aerr := receiver.AuctionRepository.Update(auction)
	if aerr != nil {
		return aerr
	}
	if err := receiver.BidRepository.Delete(newBid); err != nil {
		return err
	}

	return nil
}
