package the_number_of_beautiful_subsets

import "testing"

func Test_beautifulSubsets(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				nums: []int{2, 4, 6},
				k:    2,
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				nums: []int{18, 10, 2, 18},
				k:    8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beautifulSubsets(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("beautifulSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
