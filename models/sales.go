package models

type sales struct {
	data map[int]int
}

// TODO: instantiate from transaction model
func NewSales(d []int) Transactions {
	return NewTransactions(d)
}

func (s *sales) Get(idx int) int {
	return s.data[idx]
}
