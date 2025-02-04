package banking

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		iban string
	}
	tests := []struct {
		name    string
		args    args
		want    IBAN
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Valid IBAN", args{"GB82WEST12345698765432"}, IBAN{"GB82WEST12345698765432"}, false},
		{"Invalid IBAN", args{"GB82WEST12345698765431"}, IBAN{""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIBAN(tt.args.iban)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
