package diffence

import (
	"reflect"
	"testing"
)

func TestNewDiffer(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want *diff
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDiffer(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDiffer() = %v, want %v", got, tt.want)
			}
		})
	}
}
