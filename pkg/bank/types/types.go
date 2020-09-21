package types
type Category string
type Payment struct {
	ID  int
	Amount Money
	Category Category
}
type Money int64
type Currency string
const (
	TJS Currency = "TJS"
	RUB Currency="RUB"
	USD Currency= "USD"
)
type PAN string
type Card struct {
	ID int
	PAN PAN
	Balance Money
	Currency Currency
	Color string
	Name string
	Active bool
	MinBalance Money
}
type Result struct {
	ID int
	PAN PAN
	Balance Money
	Currency Currency
	Color string
	Name string
	Active bool
	Limit Money
}
type PaymentSource struct {
	Number PAN
	Balance Money
}

