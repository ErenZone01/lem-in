package UnitTest

import (
	"lemin/models"
	"lemin/pkg"
	"reflect"
	"testing"
)

func TestEdgeLinks(t *testing.T) {
	type args struct {
		room    string
		tunnels map[string]models.Tunnel
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.EdgeLinks(tt.args.room, tt.args.tunnels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EdgeLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeSubstrings(t *testing.T) {
	type args struct {
		input      string
		substrings []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.RemoveSubstrings(tt.args.input, tt.args.substrings); got != tt.want {
				t.Errorf("removeSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindLinks(t *testing.T) {
	type args struct {
		tunnels map[string]models.Tunnel
		rooms   map[string]models.Room
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkg.FindLinks(tt.args.tunnels, tt.args.rooms)
		})
	}
}

func TestBuildFlowNetwork(t *testing.T) {
	type args struct {
		rooms   map[string]models.Room
		tunnels map[string]models.Tunnel
	}
	tests := []struct {
		name string
		args args
		want map[string]map[string]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.BuildFlowNetwork(tt.args.rooms, tt.args.tunnels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildFlowNetwork() = %v, want %v", got, tt.want)
			}
		})
	}
}
