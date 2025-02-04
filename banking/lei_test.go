package banking

import (
	"testing"
)

func TestNewLEI(t *testing.T) {
	type args struct {
		lei string
	}
	tests := []struct {
		name    string
		args    args
		want    LEI
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Valid LEI", args{"9845000D076TY6C96A71"}, LEI{"9845000D076TY6C96A71"}, false},
		{"Invalid LEI", args{"213800A8Y1XKQMG8S713"}, LEI{""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("NewLEI(%v) = %v", tt.args.lei, tt.want)

			got, err := NewLEI(tt.args.lei)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLEI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(got.String() == tt.want.String()) {
				t.Errorf("NewLEI() = %v, want %v", got, tt.want)
			}
			t.Logf("NewLEI(%v) = %v, got %v", tt.args.lei, tt.want, got)
		})
	}
}
