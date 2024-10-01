package main

import (
	"testing"

	"github.com/JuramentoEl/GoBasic/hw06_testing/hw02_fix_app/reader"
	"github.com/JuramentoEl/GoBasic/hw06_testing/hw02_fix_app/types"
	"github.com/stretchr/testify/require"
)

func TestReadJSON(t *testing.T) {
	emp1 := types.Employee{UserID: 10, Age: 25, Name: "Rob", DepartmentID: 3}
	emp2 := types.Employee{UserID: 11, Age: 30, Name: "George", DepartmentID: 2}

	tests := []struct {
		name     string
		path     string
		expected []types.Employee
	}{
		{"data", "data.json", []types.Employee{emp1, emp2}},
		{"пустая строка", "", nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := reader.ReadJSON(tc.path)
			require.Equal(t, tc.expected, got)
		})
	}
}
