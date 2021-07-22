package arraysort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	actual := []int{6, 5, 2, 3, 9, 15, 1, 6}
	expected := []int{1, 2, 3, 5, 6, 6, 9, 15}
	sortHandler := NewSortHandler()
	tests := []struct {
		name string
		fn   func()
	}{
		{
			"test quick sort of array",
			func() {
				sortHandler.QuickSort(actual)
				if !reflect.DeepEqual(actual, expected) {
					t.Error("quick algo failed - result not as expected")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
		})
	}
}
