package domain

type BidWinnerStrategy interface {
	FindWinner(bids []*Bid) *Bid
}

type HigheshtUniqueBidStrategy struct {
}

func (s *HigheshtUniqueBidStrategy) FindWinner(bids []*Bid) *Bid {
	bidAmountMap := make(map[int][]*Bid)

	for _, bid := range bids {
		bidAmountMap[bid.BidAmount] = append(bidAmountMap[bid.BidAmount], bid)
	}

	var highestUniqueBid *Bid
	for _, bid := range bids {
		if len(bidAmountMap[bid.BidAmount]) == 1 {
			if highestUniqueBid == nil || highestUniqueBid.BidAmount < bid.BidAmount {
				highestUniqueBid = bid
			}
		}
	}

	return highestUniqueBid
}

type LowestUniqueBidStrategy struct {
}

func (s *LowestUniqueBidStrategy) FindWinner(bids []*Bid) *Bid {
	bidAmountMap := make(map[int][]*Bid)

	for _, bid := range bids {
		bidAmountMap[bid.BidAmount] = append(bidAmountMap[bid.BidAmount], bid)
	}

	var lowestUniqueBid *Bid

	for _, bid := range bids {
		if len(bidAmountMap[bid.BidAmount]) == 1 {
			if lowestUniqueBid == nil || lowestUniqueBid.BidAmount > bid.BidAmount {
				lowestUniqueBid = bid
			}
		}
	}

	return lowestUniqueBid
}
