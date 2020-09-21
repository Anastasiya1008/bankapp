package stats

import (
	"github.com/Anastasiya1008/bank/pkg/bank/types"
	//"bank/pkg/bank/types"
)
// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	summa:=types.Money(0)
	//payments=[]types.Payment{}
	for _, payment:= range payments{
		if payment.Amount>0 {
			summa+=payment.Amount
		} 
	} 
	average:=summa/types.Money(len(payments))
	return average
}