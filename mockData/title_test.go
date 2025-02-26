package mockData

import "testing"

func TestIsValidTitle(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Mr", args{"Mr"}, true},
		{"Mrs", args{"Mrs"}, true},
		{"Miss", args{"Miss"}, true},
		{"Ms", args{"Ms"}, true},
		{"Dr", args{"Dr"}, true},
		{"Prof", args{"Prof"}, true},
		{"Bum", args{"Bum"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidTitle(tt.args.in); got != tt.want {
				t.Errorf("IsValidTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
