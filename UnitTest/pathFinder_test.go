package UnitTest

import (
	"lemin/models"
	"lemin/pkg"
	"reflect"
	"testing"
	"time"
)

func TestFindPaths(t *testing.T) {
	type args struct {
		rooms   map[string]models.Room
		start   string
		end     string
		path    []string
		paths   *[][]string
		timeout time.Duration
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
			if got := pkg.FindPaths(tt.args.rooms, tt.args.start, tt.args.end, tt.args.path, tt.args.paths, tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
