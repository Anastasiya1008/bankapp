package payment
import (
	"bank/pkg/bank/types"
	//"fmt"
)
func Max(payments[] types.Payment) types.Payment {
	max:=payments[0]
	for _, payment:=range payments{
		if max.Amount<payment.Amount {
			max=payment
		}
	}	
	return max
}
func PaymentSources(cards []types.Card) []types.PaymentSource {
	for _, card:= range cards{
		if card.Balance>0 && card.Active {
			PaymentSource.Balance=card.Balance
			PaymentSource.Number=card.Number
	} 
} return PaymentSource
}