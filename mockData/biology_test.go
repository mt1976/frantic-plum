package mockData

import "testing"

func TestIsValidBiology(t *testing.T) {
	type args struct {
		biology string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"M", args{"M"}, true},
		{"F", args{"F"}, true},
		{"I", args{"I"}, true},
		{"O", args{"O"}, true},
		{"X", args{"X"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBiology(tt.args.biology); got != tt.want {
				t.Errorf("IsValidBiology() = %v, want %v", got, tt.want)
			}
		})
	}
}
