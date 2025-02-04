package banking

import (
	"testing"
)

func TestISIN_IsValid(t *testing.T) {
	tests := []struct {
		name string
		I    *ISIN
		want bool
	}{
		// TODO: Add test cases.
		{"APPLE INC", &ISIN{"US0378331005"}, true},
		{"Invalid ISIN", &ISIN{"US0378331001"}, false},
		{"WALMART", &ISIN{"US9311421039"}, true},
		{"TEST", &ISIN{"US0378331005"}, true},
		{"BAE Systems", &ISIN{"GB0002634946"}, true},
		{"Bank of Ireland", &ISIN{"6354002WOGLFYPOC1W29"}, false},
		{"Invalid Country Code", &ISIN{"XX0002634946"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.I.IsValid(); got != tt.want {
				t.Errorf("ISIN.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestISIN_Printable(t *testing.T) {
	tests := []struct {
		name string
		I    *ISIN
		want string
	}{
		// TODO: Add test cases.
		{"APPLE INC", &ISIN{"US0378331005"}, "US 037833100 5"},
		{"Invalid ISIN", &ISIN{"US0378331001"}, "US 037833100 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.I.Printable(); got != tt.want {
				t.Errorf("ISIN.Printable() = %v, want %v", got, tt.want)
			}
		})
	}
}
