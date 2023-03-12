package count_the_number_of_beautiful_subarrays

import "testing"

func Test_beautifulSubArrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1",
			args: args{nums: []int{4, 3, 1, 2, 4}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beautifulSubArrays(tt.args.nums); got != tt.want {
				t.Errorf("beautifulSubArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
