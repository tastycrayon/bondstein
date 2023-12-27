package main

import (
	"errors"
	"testing"
)

func Test_Soulution_1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		wantErr error
	}{
		{
			name: "testing current inputs",
			args: args{nums: []int{2, 3, 4}, target: 6},
			want: 2, want1: 0, wantErr: nil,
		},
		{
			name: "testing current inputs",
			args: args{nums: []int{2, 3, 5}, target: 6},
			want: 2, want1: 0, wantErr: ErrorNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Soulution_1(tt.args.nums, tt.args.target)

			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Soulution_1() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if l := len(tt.args.nums); got >= l || got1 >= l {
				t.Errorf("Soulution_1() got = %v, want %v", got1, tt.want1)
			}
			if tt.args.target != tt.args.nums[got]+tt.args.nums[got1] {
				t.Errorf("Soulution_1() got = %v %v, want %v %v", got, got1, tt.want, tt.want1)
			}
		})
	}
}
func Test_Soulution_2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		wantErr error
	}{
		{
			name: "testing current inputs",
			args: args{nums: []int{2, 3, 4}, target: 6},
			want: 2, want1: 0, wantErr: nil,
		},
		{
			name: "testing current inputs",
			args: args{nums: []int{2, 3, 5}, target: 6},
			want: 2, want1: 0, wantErr: ErrorNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Soulution_2(tt.args.nums, tt.args.target)

			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Soulution_2() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if l := len(tt.args.nums); got >= l || got1 >= l {
				t.Errorf("Soulution_2() got = %v, want %v", got1, tt.want1)
			}
			if tt.args.target != tt.args.nums[got]+tt.args.nums[got1] {
				t.Errorf("Soulution_2() got = %v %v, want %v %v", got, got1, tt.want, tt.want1)
			}
		})
	}
}
