package filter

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test int",
			args: struct {
				slice1 []int
				slice2 []int
			}{slice1: []int{1, 3, 4, 5, 7}, slice2: []int{1, 3, 4, 6, 5}},
			want: []int{1, 3, 4, 5, 7, 6},
		},
		{
			name: "test string",
			args: struct {
				slice1 []int
				slice2 []int
			}{slice1: []int{1, 3, 4, 5}, slice2: []int{1, 3, 4, 5}},
			want: []int{1, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
