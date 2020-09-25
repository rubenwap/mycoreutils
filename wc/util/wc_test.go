package util

import (
	"reflect"
	"testing"

	"github.com/urfave/cli"
)

func TestWc(t *testing.T) {
	tests := []struct {
		name string
		want *cli.App
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wc() = %v, want %v", got, tt.want)
			}
		})
	}
}
