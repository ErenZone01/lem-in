package UnitTest

import (
	"lemin/models"
	"lemin/pkg"
	"testing"
)

func TestMove(t *testing.T) {
	type args struct {
		paths [][]models.Room
		ant   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkg.Move(tt.args.paths, tt.args.ant)
		})
	}
}

func Test_release(t *testing.T) {
	type args struct {
		paths [][]models.Room
		tab   map[int][]string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkg.Release(tt.args.paths, tt.args.tab)
		})
	}
}

func Test_sortListByNumber(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkg.SortListByNumber(tt.args.list)
		})
	}
}

func Test_extractNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.ExtractNumber(tt.args.s); got != tt.want {
				t.Errorf("extractNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckStringInArray(t *testing.T) {
	type args struct {
		arr    []string
		target string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.CheckStringInArray(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("CheckStringInArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
