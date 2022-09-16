package strs

import (
	"testing"
)

func TestNextAsciiStr1(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		// TODO: Add test cases.
		{name: "", args: args{c: 'a'}, want: 'b'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextAsciiStr(tt.args.c); got != tt.want {
				t.Errorf("NextAsciiStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
