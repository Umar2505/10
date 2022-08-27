package stats

import "github.com/Umar2505/bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	var result types.Money
	for _,el:= range payments {
		result+=el.Amount
	}
	result/=types.Money(len(payments))
	return result
}