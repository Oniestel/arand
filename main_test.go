package main

import "testing"

func Test_randInt(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name    string
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name:    "Test max = 10",
			args:    args{10},
			wantN:   10,
			wantErr: false,
		},
		{
			name:    "Test max = 1000000000",
			args:    args{1000000000},
			wantN:   1000000000,
			wantErr: false,
		},
		{
			name:    "Test max = 0",
			args:    args{0},
			wantN:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := randInt(tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("randInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN > tt.wantN {
				t.Errorf("randInt() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
