package financial

import (
	"reflect"
	"testing"
)

func TestNewTenor(t *testing.T) {
	type args struct {
		term string
	}

	tests := []struct {
		name    string
		args    args
		want    Tenor
		wantErr bool
	}{
		// TODO: Add test cases.
		{"SP", args{"SP"}, Tenor{"SP"}, false},
		{"td", args{"td"}, Tenor{"TD"}, false},
		{"1D", args{"1D"}, Tenor{"1D"}, false},
		{"1W", args{"1W"}, Tenor{"1W"}, false},
		{"POO", args{"POO"}, Tenor{""}, true},
		{"1X", args{"1X"}, Tenor{""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTenor(tt.args.term)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTenor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTenor() = %v, want %v", got, tt.want)
			}
			t.Logf("NewTenor() = sent[%v], returned [%v], want [%v], err [%v]", tt.args.term, got, tt.want, err)
		})
	}
}
