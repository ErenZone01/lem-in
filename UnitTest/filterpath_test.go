package UnitTest

import (
	"lemin/models"
	"lemin/pkg"
	"reflect"
	"testing"
)

func TestIsOptimal(t *testing.T) {
	type args struct {
		path1 []string
		path2 []string
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
			if got := pkg.IsOptimal(tt.args.path1, tt.args.path2); got != tt.want {
				t.Errorf("IsOptimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minmax(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.Minmax(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minmax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBestPath(t *testing.T) {
	type args struct {
		PathGroups [][][]string
		ants       int
	}
	tests := []struct {
		name string
		args args
		want [][]models.Room
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.BestPath(tt.args.PathGroups, tt.args.ants); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		ant   int
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want [][][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.Filter(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allPathCollision(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want [][][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.AllPathCollision(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
