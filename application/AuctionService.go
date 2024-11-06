package application

import (
	"errors"
	"super-bidder/domain"
	"super-bidder/infra"
)

type AuctionService struct {
	domain.AuctionRepository
	domain.SellerRepository
}

func NewAuctionService(sellerRepo domain.SellerRepository) *AuctionService {
	return &AuctionService{AuctionRepository: infra.NewAuctionRepository(),
		SellerRepository: sellerRepo}
}

func (receiver *AuctionService) Create(sellerName string, minA int, maxA int, aId string) error {
	if seller, err := receiver.SellerRepository.Get(sellerName); err != nil {
		return err
	} else {
		auction := seller.Auctions(minA, maxA, aId)
		if err := receiver.AuctionRepository.Create(auction); err != nil {
			return err
		}
	}

	return nil
}

func (receiver *AuctionService) Close(aId string) (error, *domain.Bid) {
	if auction, err := receiver.AuctionRepository.Get(aId); err != nil {
		return err, nil
	} else {
		bid := auction.Close(&domain.HigheshtUniqueBidStrategy{})

		if bid != nil {
			bid.Buyer.UpdateBalance(bid.Buyer.Balance() - bid.BidAmount)
			// update buyer
			if _, err := infra.NewBuyerRepository().Update(bid.Buyer); err != nil {
				return err, nil
			}
			return nil, bid
		} else {
			return errors.New("no bids found"), nil
		}
	}
}
