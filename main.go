package main

import (
	"fmt"
	"super-bidder/application"
	"super-bidder/infra"
)

func main() {
	sr := infra.NewSellerRepository()
	br := infra.NewBuyerRepository()
	ar := infra.NewAuctionRepository()
	bir := infra.NewBidRepository()

	bs := application.BuyerService{
		BuyerRepository: br,
	}
	if err := bs.AddBuyer("buyer1", 20); err != nil {
		fmt.Println(err)
	}
	if err := bs.AddBuyer("buyer2", 20); err != nil {
		fmt.Println(err)
	}
	if err := bs.AddBuyer("buyer3", 20); err != nil {
		fmt.Println(err)
	}

	ss := application.SellerService{
		SellerRepository: sr,
	}
	if err := ss.AddSeller("seller1"); err != nil {
		fmt.Println(err)
	}
	as := application.AuctionService{
		SellerRepository:  sr,
		AuctionRepository: ar,
	}
	err := as.Create("seller1", 10, 50, "A1")
	if err != nil {
		fmt.Println(err)
	}

	bid := application.NewBidService(bir, ar, br)
	erra := bid.CreateBid("buyer1", "A1", 17)
	if erra != nil {
		fmt.Println(erra)
	}
	if err := bid.CreateBid("buyer2", "A1", 19); err != nil {
		fmt.Println(err)
	}
	if err := bid.CreateBid("buyer3", "A1", 19); err != nil {
		fmt.Println(err)
	}
	//	CloseAuction
	if err, bid := as.Close("A1"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bid.Buyer.Name())
	}

	//	Testcase 2
	if err := ss.AddSeller("seller2"); err != nil {
		fmt.Println(err)
	}

	if err := as.Create("seller2", 5, 20, "A2"); err != nil {
		fmt.Println(err)
	}

	if err := bid.CreateBid("buyer3", "A2", 25); err != nil {
		fmt.Println(err)
	}
	if err := bid.CreateBid("buyer2", "A2", 5); err != nil {
		fmt.Println(err)
	}
	if err := bid.Withdraw("buyer2", "A2"); err != nil {
		fmt.Println(err)
	}
	if err, bid := as.Close("A2"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bid.Buyer.Name())
	}
}
