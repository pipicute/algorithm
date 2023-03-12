package utils

import (
	"reflect"
	"testing"
)

func TestGetPrimes(t *testing.T) {
	type args[T interface{ int | int64 }] struct {
		maxNum T
	}
	type testCase[T interface{ int | int64 }] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "1",
			args: args[int]{maxNum: 97},
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPrimes(tt.args.maxNum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortIndex(t *testing.T) {
	type args[T interface{ int | int64 }] struct {
		l   []T
		asc bool
	}
	type testCase[T interface{ int | int64 }] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "1",
			args: args[int]{l: []int{-1, -1, 3, 2}, asc: false},
			want: []int{2, 3, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortIndex(tt.args.l, tt.args.asc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	type args[T interface{ int | int64 }] struct {
		num T
	}
	type testCase[T interface{ int | int64 }] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "1",
			args: args[int]{19},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.num); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
