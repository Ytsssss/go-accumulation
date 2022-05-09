package utils

import (
	"reflect"
	"testing"
)

func TestDiffStringSlice(t *testing.T) {
	type args struct {
		newElems []string
		oldElems []string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		{
			name: "success",
			args: args{
				newElems: []string{
					"123", "234", "345",
				},
				oldElems: []string{
					"236", "345", "111",
				},
			},
			want: []string{
				"123", "234",
			},
			want1: []string{
				"236", "111",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := DiffStringSlice(tt.args.newElems, tt.args.oldElems)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiffStringSlice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DiffStringSlice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
