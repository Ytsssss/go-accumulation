package utils

import (
	"log"
	"testing"
)

func TestSafeGo(t *testing.T) {
	type args struct {
		f func()
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				f: func() {
					i := 10
					log.Printf("i=%d", i)
				},
			},
		}, {
			name: "panic",
			args: args{
				f: func() {
					panic("panic error")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SafeGo(tt.args.f)
		})
	}
}
