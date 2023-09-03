package main

import (
	"reflect"
	"testing"
)

func Test_getConcatenation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{nums: []int{1, 2, 3}},
			want: []int{1, 2, 3, 1, 2, 3},
		},
		{
			name: "example 2",
			args: args{nums: []int{3, 4, 7}},
			want: []int{3, 4, 7, 3, 4, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConcatenation(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConcatenation() = %v, want %v", got, tt.want)
			}
		})
	}
}
