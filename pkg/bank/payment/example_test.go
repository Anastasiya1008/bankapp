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