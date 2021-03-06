package card
import (
	"github.com/Anastasiya1008/bankapp/pkg/bank/types"
	"fmt"
	)

func ExampleWithdraw_positive() {
card:=types.Card{Balance:20_000_00, Active: true}
Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 1000000
}
func ExampleWithdraw_noMoney(){
	card:=types.Card{Balance:0_00, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)

	// Output: 0
}
func ExampleWithdraw_inactive(){
	card:=types.Card{Balance:20_000_00, Active: false}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)

	// Output: 2000000
}
func ExampleWithdraw_limit(){
	card:=types.Card{Balance:20_000_00, Active: true}
	Withdraw(&card, 30_000_00)
	fmt.Println(card.Balance)

	// Output: 2000000
}
func ExampleDeposit_positive(){
	card:=types.Card{Balance:20_000_00, Active: true}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)

	// Output: 3000000
}
func ExampleDeposit_negativeBalance(){
	card:=types.Card{Balance:-20_000_00, Active: true}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)

	// Output: -1000000
}
func ExampleDeposit_notActive(){
	card:=types.Card{Balance:20_000_00, Active: false}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)

	// Output: 2000000
}

func ExampleDeposit_aboveLimit(){
	card:=types.Card{Balance:20_000_00, Active: true}
	Deposit(&card, 60_000_00)
	fmt.Println(card.Balance)

	// Output: 2000000
}
func ExampleAddBonus_positive(){
	card:=types.Card{Balance:10_000_00, MinBalance:10_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	// Output: 1002465
}
func ExampleAddBonus_notactive(){
	card:=types.Card{Balance:10_000_00, MinBalance:10_000_00, Active: false}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	// Output: 1000000
}
func ExampleAddBonus_negative(){
	card:=types.Card{Balance:-10_000_00, MinBalance:10_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	// Output: -1000000
}
func ExampleAddBonus_overbonus(){
	card:=types.Card{Balance:4_000_000_00,MinBalance:3_000_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)

	// Output: 400000000
}
func ExampleTotal () {
	cards:=[]types.Card {
		types.Card {
			Balance: 10_000_00,
			Active: true,
		},
		types.Card {
			Balance: 10_000_00,
			Active: true,
		},
	}
	fmt.Println(Total(cards))
	// Output: 2000000
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

