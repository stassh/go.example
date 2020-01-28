package main

import "testing"

func TestSum(t *testing.T) {

	testTables := []struct {
		inputA int64
		inputB int64
		expected int64
	} {
		{inputA: 1, inputB: 1, expected: 2},
		{inputA: 1, inputB: 2, expected: 3},
		{inputA: 2, inputB: 3, expected: 5},
	}

	for _, testTable := range testTables {
		actualResult := Sum(testTable.inputA, testTable.inputB)
		if actualResult != testTable.expected {
			t.Errorf("\nExpected: %v, Actual: %v", testTable.expected, actualResult)
		}
	}
}