package payment
import (
	"github.com/Anastasiya1008/bank/pkg/bank/types"
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
	fmt.Println(max.ID)
	// Output: 5
}
func ExamplePaymentSources() {
	card:=[]types.Card {
		{
			PAN:"12",
			Balance:-20_000_00,
		},
		{ 	
			PAN: "13",
			Balance:90_000_00,
		},
		{ 	
			Active: true,
			PAN:"45",
			Balance:23_000_00,
		},
		{ 	
			Active: true,
			PAN:"54",
			Balance:32_000_00,
		},
	}
	payments:=PaymentSources(card)
	fmt.Println(payments)
	// Output: [{45 2300000} {54 3200000}]
}	
