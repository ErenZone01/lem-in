package UnitTest

import (
	"lemin/models"
	"lemin/pkg"
	"reflect"
	"testing"
)

func Test_isValidRoom(t *testing.T) {
	type args struct {
		line string
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
			if got := pkg.IsValidRoom(tt.args.line); got != tt.want {
				t.Errorf("isValidRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isComment(t *testing.T) {
	type args struct {
		element string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Define test cases
		{name: "test case 1", args: args{element: "# Comment"}, want: true},
		{name: "test case 2", args: args{element: "Not a comment"}, want: false},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.IsComment(tt.args.element); got != tt.want {
				t.Errorf("isComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isStart(t *testing.T) {
	type args struct {
		element string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test case 1 ", args: args{element: "##start"}, want: true},
		{name: "test case 2 ", args: args{element: "##end"}, want: false},
		{name: "test case 3 ", args: args{element: "#end"}, want: false},
		{name: "test case 3 ", args: args{element: "#start"}, want: false},
		{name: "test case 4 ", args: args{element: "start"}, want: false},
		{name: "test case 5 ", args: args{element: "##start   "}, want: false},
		{name: "test case 6 ", args: args{element: "##start 1 2"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.IsStart(tt.args.element); got != tt.want {
				t.Errorf("isStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEnd(t *testing.T) {
	type args struct {
		element string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test case 1 ", args: args{element: "##end"}, want: true},
		{name: "test case 2 ", args: args{element: "#end"}, want: false},
		{name: "test case 3 ", args: args{element: "#start"}, want: false},
		{name: "test case 3 ", args: args{element: "##start"}, want: false},
		{name: "test case 4 ", args: args{element: "end"}, want: false},
		{name: "test case 5 ", args: args{element: "##end   "}, want: false},
		{name: "test case 6 ", args: args{element: "##end 1 2"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.IsEnd(tt.args.element); got != tt.want {
				t.Errorf("isEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPositiveInteger(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test case 1 ", args: args{s: "3"}, want: true},
		{name: "test case 2 ", args: args{s: "0"}, want: true},
		{name: "test case 3 ", args: args{s: "-1"}, want: true},
		{name: "test case 4 ", args: args{s: "1fg"}, want: false},
		{name: "test case 5 ", args: args{s: "2.3"}, want: false},
		{name: "test case 6 ", args: args{s: "  "}, want: false},
		{name: "test case 7 ", args: args{s: "#start"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.IsPositiveInteger(tt.args.s); got != tt.want {
				t.Errorf("isPositiveInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFloat(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test case 1 ", args: args{s: "3.5"}, want: true},
		{name: "test case 2 ", args: args{s: "0"}, want: true},
		{name: "test case 3 ", args: args{s: "-1"}, want: true},
		{name: "test case 4 ", args: args{s: "1fg"}, want: false},
		{name: "test case 5 ", args: args{s: "2.3"}, want: true},
		{name: "test case 6 ", args: args{s: "  "}, want: false},
		{name: "test case 7 ", args: args{s: "#start"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.IsFloat(tt.args.s); got != tt.want {
				t.Errorf("isFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidTunnel(t *testing.T) {
	type args struct {
		tunnel string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test case 1", args: args{tunnel: "1-3"}, want: true},
		{name: "test case 2", args: args{tunnel: "13"}, want: false},
		{name: "test case 3", args: args{tunnel: "1|3"}, want: false},
		{name: "test case 4", args: args{tunnel: "1*3"}, want: false},
		{name: "test case 5", args: args{tunnel: "1+3"}, want: false},
		{name: "test case 6", args: args{tunnel: "1n3"}, want: false},
		{name: "test case 7", args: args{tunnel: "1.3"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.IsValidTunnel(tt.args.tunnel); got != tt.want {
				t.Errorf("isValidTunnel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasDuplicateRoom(t *testing.T) {
	type args struct {
		rooms map[string]models.Room
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test case 1", args: args{rooms: map[string]models.Room{}}, want: false},
		{name: "test case 2", args: args{rooms: map[string]models.Room{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkg.HasDuplicateRoom(tt.args.rooms); got != tt.want {
				t.Errorf("HasDuplicateRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataParser(t *testing.T) {
	type args struct {
		path    string
		rooms   map[string]models.Room
		tunnels map[string]models.Tunnel
	}
	tests := []struct {
		name           string
		args           args
		wantNumants    string
		wantStartPoint string
		wantEndPoint   string
		wantInput      []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNumants, gotStartPoint, gotEndPoint, gotInput := pkg.DataParser(tt.args.path, tt.args.rooms, tt.args.tunnels)
			if gotNumants != tt.wantNumants {
				t.Errorf("DataParser() gotNumants = %v, want %v", gotNumants, tt.wantNumants)
			}
			if gotStartPoint != tt.wantStartPoint {
				t.Errorf("DataParser() gotStartPoint = %v, want %v", gotStartPoint, tt.wantStartPoint)
			}
			if gotEndPoint != tt.wantEndPoint {
				t.Errorf("DataParser() gotEndPoint = %v, want %v", gotEndPoint, tt.wantEndPoint)
			}
			if !reflect.DeepEqual(gotInput, tt.wantInput) {
				t.Errorf("DataParser() gotInput = %v, want %v", gotInput, tt.wantInput)
			}
		})
	}
}

func TestFindUnknownRoom(t *testing.T) {
	type args struct {
		tunnels map[string]models.Tunnel
		rooms   map[string]models.Room
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
			if got := pkg.FindUnknownRoom(tt.args.tunnels, tt.args.rooms); got != tt.want {
				t.Errorf("FindUnknownRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataValidation(t *testing.T) {
	type args struct {
		start      string
		end        string
		startLinks []string
		endLinks   []string
		rooms      map[string]models.Room
		tunnels    map[string]models.Tunnel
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkg.DataValidation(tt.args.start, tt.args.end, tt.args.startLinks, tt.args.endLinks, tt.args.rooms, tt.args.tunnels)
		})
	}
}
