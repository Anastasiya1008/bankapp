package payment
import (
	"bank/pkg/bank/types"
	"fmt"

)
func ExampleMax(){
	payment:=[]types.Payment {
		{
			ID:12,
			Amount:20_000_00,
		},
		{
			ID:34,
			Amount:10_000_00,
		},
		{ 	
			ID:5,
			Amount:90_000_00,
		},
		{ 	
			ID:45,
			Amount:23_000_00,
		},
	}
	max:=Max(payment)
	fmt.Println(max)
	// Output: {5 9000000}
}
func ExamplePaymentSources_() {
	card:=[]types.Card {
		{
			Number:12,
			Balance:20_000_00,
		},
		{
			Number:34,
			Balance:-10_000_00,
		},
		{ 	
			Active: false,
			Amount:90_000_00,
		},
		{ 	
			Number:45,
			Amount:23_000_00,
		},
	}
	payment=PaymentSources(cards)
	fmt.Println(payment)
	// Output: {}
}	
