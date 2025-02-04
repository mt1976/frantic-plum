package financial

import (
	"reflect"
	"testing"
	"time"
)

func Test_validateAndFormatTerm(t *testing.T) {
	type args struct {
		term string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"SP", args{"SP"}, "SP", false},
		{"td", args{"td"}, "TD", false},
		{"1D", args{"1D"}, "1D", false},
		{"1W", args{"1W"}, "1W", false},
		{"1M", args{"1M"}, "1M", false},
		{"1Y", args{"1Y"}, "1Y", false},
		{"1d", args{"1d"}, "1D", false},
		{"10w", args{"10w"}, "10W", false},
		{"10m", args{"10m"}, "10M", false},
		{"10X", args{"10X"}, "", true},
		{"BUMBUM", args{"BUMBUM"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateAndFormatTenor(tt.args.term)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateAndFormatTerm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validateAndFormatTerm() = %v, want %v, %e", got, tt.want, err)
			}
			t.Logf("validateAndFormatTerm() = %v, want %v, error %v", got, tt.want, err)
		})
	}
}

func Test_getTenorDate(t *testing.T) {
	type args struct {
		tenor     Tenor
		tradeDate time.Time
		ccy       []string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{"SPUSD", args{Tenor{"SP"}, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), []string{"USD"}}, time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{"TDUSD", args{Tenor{"TD"}, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), []string{"USD"}}, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), false},
		{"SPGBPUSD", args{Tenor{"SP"}, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), []string{"GBP", "USD"}}, time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), false},
		{"SPEURZARvUSD", args{Tenor{"SP"}, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), []string{"EUR", "ZAR", "USD"}}, time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), false},
		{"SPMXN", args{Tenor{"SP"}, time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), []string{"MXN"}}, time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), false},
		{"SPMXNvUSD", args{Tenor{"SP"}, time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), []string{"MXN", "USD"}}, time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), false},
		{"SPMXNvUSDEUR", args{Tenor{"SP"}, time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), []string{"MXN", "USD", "EUR"}}, time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), false},
		{"TDGBP", args{Tenor{"TD"}, time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), []string{"GBP"}}, time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{"1MGBPvUSD", args{Tenor{"1M"}, time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), []string{"GBP", "USD"}}, time.Date(2019, 2, 3, 0, 0, 0, 0, time.UTC), false},
		{"1YGBPvUSD", args{Tenor{"1Y"}, time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), []string{"GBP", "USD"}}, time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC), false},
		{"1YUSDvCAD", args{Tenor{"1Y"}, time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), []string{"USD", "CAD"}}, time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), false},
		//{"1YUSDvCADvEUR", args{Tenor{""}, time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), []string{"USD", "CAD", "EUR"}}, time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDateFromTenor(tt.args.tenor, tt.args.tradeDate, tt.args.ccy...)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTenorDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTenorDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTenorFromDate(t *testing.T) {
	type args struct {
		inDate   time.Time
		baseDate time.Time
		ccy      []string
	}
	tests := []struct {
		name    string
		args    args
		want    Tenor
		wantErr bool
	}{
		// TODO: Add test cases.
		{"SPGBP", args{time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), []string{"GBP"}}, Tenor{"SP"}, false},
		{"1WGBP", args{time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), []string{"GBP"}}, Tenor{"1W"}, false},
		{"1WGBP", args{time.Date(2019, 1, 11, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), []string{"GBP"}}, Tenor{"1W"}, false},
		{"SPUSD", args{time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), []string{"USD"}}, Tenor{"SP"}, false},
		{"1WUSDSP", args{time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), []string{"USD"}}, Tenor{"SP"}, false},
		{"1WUSD", args{time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), []string{"USD"}}, Tenor{"1W"}, false},
		{"TDEUR", args{time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), []string{"EUR"}}, Tenor{"TD"}, false},
		{"10YUSD", args{time.Date(2033, 7, 19, 0, 0, 0, 0, time.UTC), time.Date(2023, 7, 8, 0, 0, 0, 0, time.UTC), []string{"USD"}}, Tenor{"10Y"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTenorFromDate(tt.args.inDate, tt.args.baseDate, tt.args.ccy...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTenorFromDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTenorFromDate() = %v, want %v", got, tt.want)
			}
			t.Logf("TEST >>>> Want: %v, Got: %v, WantErr %v\n", tt.want, got, tt.wantErr)
		})
	}
}

func TestGetLadder(t *testing.T) {
	type args struct {
		pivotDate time.Time
		ccy       []string
	}
	tests := []struct {
		name string
		args args
		//want    []FinDate
		//want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"1YUSD", args{time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), []string{"USD"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2, err := GetLadder(tt.args.pivotDate, tt.args.ccy...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLadder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("TEST >>>> Want: %v, Got: %v, WantErr %v\n", tt.wantErr, got, got2)

		})
	}
}
