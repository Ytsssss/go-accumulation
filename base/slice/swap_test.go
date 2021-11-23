package slice

import "testing"

func Test_swapError1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swapError1()
		})
	}
}

func Test_swapError2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swapError2()
		})
	}
}
