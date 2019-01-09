package main

import (
	"reflect"
	"testing"
)

func TestArrange(t *testing.T) {
	type args struct {
		repeat     int
		operations [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Example#1", args: args{repeat: 2, operations: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}}}, want: []int{1, 4, 2, 3, 5, 6, 7, 8}},
		{name: "Example#2", args: args{repeat: 1000000000, operations: [][]int{{1, 3}, {6, 8}, {3, 5}, {2, 6}, {3, 7}, {3, 4}, {4, 7}, {2, 4}, {1, 3}, {2, 7}, {2, 7}, {2, 4}, {6, 7}, {1, 7}, {3, 4}, {1, 6}}}, want: []int{1, 8, 3, 4, 5, 2, 7, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Arrange(tt.args.repeat, tt.args.operations); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Arrange() = %v, want %v", got, tt.want)
			}
		})
	}
}
