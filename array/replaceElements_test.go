package main

import (
	"reflect"
	"testing"
)

func Test_replaceElements(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{
			name: "example 1",
			args: args{[]int{17, 18, 5, 4, 6, 1}},
			want: []int{18, 6, 6, 6, 1, -1},
		},
		{
			name: "example 2",
			args: args{[]int{400}},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceElements(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceElements() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceElements2(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceElements2() = %v, want %v", got, tt.want)
			}
		})
	}
}
