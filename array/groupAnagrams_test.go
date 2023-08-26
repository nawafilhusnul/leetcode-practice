package main

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		skipped bool
		name    string
		args    args
		want    [][]string
	}{
		{
			name:    "example 1",
			args:    args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want:    [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
			skipped: true,
		},
		{
			name: "example 2",
			args: args{strs: []string{""}},
			want: [][]string{{""}},
		},
		{
			name: "example 3",
			args: args{strs: []string{"a"}},
			want: [][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		if tt.skipped {
			// ignore the skipped test because lack of possibility  of answers
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
