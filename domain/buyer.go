package domain

type Buyer interface {
	Bids(auctionName string, amount int)
	UniqueId() string
	Balance() int
	UpdateBalance(b int)
	Name() string
}

type DefaultBuyer struct {
	name      string
	preferred bool
	budget    int
}

func (b *DefaultBuyer) Name() string {
	return b.name
}

func NewBuyer(name string, budget int) Buyer {
	return &DefaultBuyer{preferred: false,
		name:   name,
		budget: budget}
}

func (b *DefaultBuyer) Bids(auctionName string, amount int) {

}

func (b *DefaultBuyer) UniqueId() string {
	return b.name
}

func (b *DefaultBuyer) Balance() int {
	return b.budget
}

func (b *DefaultBuyer) UpdateBalance(balance int) {
	b.budget = balance
}
