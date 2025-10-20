package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0
	result := CalculateTax(amount)
	
	if result != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct { 
		amount, expect float64
	}
	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}
	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %.2f, but got %.2f", item.expect, result)
		}
	}
}