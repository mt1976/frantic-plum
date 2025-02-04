package financial

import (
	"reflect"
	"testing"
	"time"
)

func Test_settlementDays(t *testing.T) {
	type args struct {
		ccy1 string
		ccy2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"EURUSD", args{"EUR", "USD"}, 2},
		{"EURGBP", args{"EUR", "GBP"}, 2},
		{"USDCAD", args{"USD", "CAD"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSettlementDaysPAIR(tt.args.ccy1, tt.args.ccy2)
			if err != nil {
				t.Errorf("settlementDays() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("settlementDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_settlementDaysVia(t *testing.T) {
	type args struct {
		ccy1 string
		via  string
		ccy2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"EURZAR via USD", args{"EUR", "USD", "ZAR"}, 2},
		{"EURMXN via USD", args{"EUR", "USD", "MXN"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSettlementDaysCROSS(tt.args.ccy1, tt.args.via, tt.args.ccy2)
			if err != nil {
				t.Errorf("settlementDaysVia() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("settlementDaysVia() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_settlementDate(t *testing.T) {
	type args struct {
		ccy1   string
		ccy2   string
		inDate time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
		{"EURUSD", args{"EUR", "USD", time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)}, time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), false},
		{"EURMXN", args{"EUR", "MXN", time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)}, time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC), false},
		{"USDCAD", args{"USD", "CAD", time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)}, time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), false},
		{"GBPUSD", args{"GBP", "USD", time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)}, time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), false},
		{"BTCETH", args{"BTC", "ETH", time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)}, time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSettlementDatePAIR(tt.args.ccy1, tt.args.ccy2, tt.args.inDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("settlementDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("settlementDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSettlementDateCCY(t *testing.T) {
	type args struct {
		ccy1   string
		inDate time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
		{"EUR", args{"EUR", time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)}, time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSettlementDateCCY(tt.args.ccy1, tt.args.inDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSettlementDateCCY() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSettlementDateCCY() = %v, want %v", got, tt.want)
			}
		})
	}
}
