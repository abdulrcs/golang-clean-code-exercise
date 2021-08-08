package models

type transactions struct {
	data map[int]int
}

// TODO: implement this
// i.e. go test is green
func NewTransactions(d []int) Transactions {
	data := make(map[int]int)

	for i, v := range d {
		data[i] = v
	}

	return &transactions{data: data}
}

// TODO: implement this
// i.e. go test is green
func (t *transactions) Get(idx int) int {
	return t.data[idx]
}

// TODO: implement this method
// i.e. go test is green
func (t *transactions) GetTotal() int {
	sum := 0
	for _, v := range t.data {
		sum += v
	}
	return sum
}

// TODO: implement this method
// i.e. go test is green
func (t *transactions) GetTotalWithinRange(i, j int) int {
	sum := 0
	for i <= j {
		sum += t.data[i]
		i++
	}
	return sum
}
