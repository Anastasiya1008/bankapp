package stats
import (
	"fmt"
	"github.com/Anastasiya1008/bank/pkg/bank/types"
)
func ExampleAvg(){
	payments:=[]types.Payment {
	{
		ID:23,
		Amount: 24_00,
		Category: "white",
	},
	
	{	ID:10,
		Amount: 10_00,
		Category: "gold",
	},
	{	ID:15,
		Amount: 64_00,
		Category: "gold",
	},
	{	ID:5,
		Amount: 66_00,
		Category: "gold",
	},
}
	Average:=Avg(payments)
	fmt.Println(Average)
// Output:4100
}