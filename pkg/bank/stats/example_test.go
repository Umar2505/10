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

func ExampleTotalInCategory_positive() {
	payments := []types.Payment{
		{Amount: 300, Category: "Online Shops"},{Amount: 150,Category: "Clothes"},{Amount: 450, Category: "Online Shops"},
	}
	result := TotalInCategory(payments,"Online Shops")
	fmt.Println(result)
	
	//Output:
	//750
	
}