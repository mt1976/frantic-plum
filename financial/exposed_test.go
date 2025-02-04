package financial

import "testing"

func TestFormatAmount(t *testing.T) {
	type args struct {
		inAmount float64
		inCCY    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"EUR", args{1234.56, "EUR"}, "1,234.56"},
		{"USD", args{1234.56, "USD"}, "1,234.56"},
		{"JPY", args{1234.56, "JPY"}, "1,235"},
		{"DASH", args{1234.56, "DASH"}, "1,234.56000000"},
		{"XAG", args{1234.56, "XAG"}, "1,235"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatAmount(tt.args.inAmount, tt.args.inCCY); got != tt.want {
				t.Errorf("FormatAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
