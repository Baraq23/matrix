package main

import (
	"testing"
)

func TestMatrixStr(t *testing.T) {
	tests := []struct {
		input string
		want  [4]int64
	}{
		{"(1,2)(3,4)", [4]int64{1, 2, 3, 4}},
		{"(5,6)(7,8)", [4]int64{5, 6, 7, 8}},
		{"(-1,2)(-3,4)", [4]int64{-1, 2, -3, 4}},
	}

	for _, tt := range tests {
		got1, got2, got3, got4 := MatrixStr(tt.input)
		got := [4]int64{got1, got2, got3, got4}
		if got != tt.want {
			t.Errorf("MatrixStr(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestMultMatrix(t *testing.T) {
	tests := []struct {
		a, b, c, d, e, f, g, h int64
		want                   [4]int64
	}{
		{1, 2, 3, 4, 5, 6, 7, 8, [4]int64{19, 22, 43, 50}},
		{2, 0, 1, 2, 1, 2, 3, 4, [4]int64{2, 4, 7, 10}},
		{-1, 2, -3, 4, -5, 6, -7, 8, [4]int64{-9, 10, -43, 50}},
	}

	for _, tt := range tests {
		got1, got2, got3, got4 := MultMatrix(tt.a, tt.b, tt.c, tt.d, tt.e, tt.f, tt.g, tt.h)
		got := [4]int64{got1, got2, got3, got4}
		if got != tt.want {
			t.Errorf("MultMatrix(%d, %d, %d, %d, %d, %d, %d, %d) = %v, want %v", tt.a, tt.b, tt.c, tt.d, tt.e, tt.f, tt.g, tt.h, got, tt.want)
		}
	}
}
