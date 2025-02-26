package mockData

import "testing"

func TestIsValidGender(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Cis", args{"Cis"}, true},
		{"Trans", args{"Trans"}, true},
		{"Bum", args{"Bum"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidGender(tt.args.in); got != tt.want {
				t.Errorf("IsValidGender() = %v, want %v", got, tt.want)
			}
		})
	}
}
