package mockData

import (
	"reflect"
	"testing"
)

func TestGetCurrencyInfo(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		args    args
		want    Currency
		wantErr bool
	}{
		{"Valid Currency Code", args{"EUR"}, Currency{Code: "EUR", SpotDays: 2, Name: "Euro", Character: "€", DPS: 2, QuoteDPS: 4, Type: Fiat, MajorUnit: "Euro", MinorUnit: "Cent", ISONumericCode: "978", KnownAs: "Euros", MinorCharacter: "c", YearOfIntroduction: 1999}, false},
		{"Invalid Currency Code", args{"XX"}, Currency{}, true},
		{"Sterling", args{"GBP"}, Currency{Code: "GBP", SpotDays: 2, Name: "Pound Sterling", Character: "£", DPS: 2, QuoteDPS: 4, Type: Fiat, MajorUnit: "Pound", MinorUnit: "Pence", ISONumericCode: "826", KnownAs: "Quids", MinorCharacter: "p", YearOfIntroduction: 800}, false},
		{"Valid Currency Code", args{"USD"}, Currency{Code: "USD", SpotDays: 1, Name: "US Dollar", Character: "$", DPS: 2, QuoteDPS: 4, Type: Fiat, MajorUnit: "Dollar", MinorUnit: "Cent", ISONumericCode: "840", KnownAs: "Bucks", MinorCharacter: "c", YearOfIntroduction: 1792}, false},
		{"Crypto DASH", args{"DASH"}, Currency{Code: "DASH", SpotDays: 0, Name: "Dash", Character: "Đ", DPS: 8, QuoteDPS: 8, Type: Crypto, MajorUnit: "Dash", MinorUnit: "Duffs", ISONumericCode: "1005", KnownAs: "Dash", MinorCharacter: "d", YearOfIntroduction: 2014}, false},
		{"Metal XAU", args{"XAU"}, Currency{Code: "XAU", SpotDays: 2, Name: "Gold", Character: "Au", DPS: 0, QuoteDPS: 0, Type: Metals, MajorUnit: "Ounce", MinorUnit: "Ounce", ISONumericCode: "959", KnownAs: "Gold"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCurrency(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}
