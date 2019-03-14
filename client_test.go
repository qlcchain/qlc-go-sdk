package qlcchain

import (
	"reflect"
	"testing"
)

func TestNewQLCClient(t *testing.T) {
	type args struct {
		endPoint string
	}
	tests := []struct {
		name    string
		args    args
		want    *QLCClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewQLCClient(tt.args.endPoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewQLCClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQLCClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
