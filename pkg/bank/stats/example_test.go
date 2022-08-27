package stats

import (
	"fmt"

	"github.com/Umar2505/bank/pkg/bank/types"
)

func ExampleAvg_positive() {
	payments := []types.Payment{
		{Amount: 300},{Amount: 150},{Amount: 450},
	}
	result := Avg(payments)
	fmt.Println(result)

	//Output:
	//300

}