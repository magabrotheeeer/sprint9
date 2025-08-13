package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{"0 size", 0, 0},
		{"small fixed size", 5, 5},
		{"large fixed size", 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := generateRandomElements(tt.size)
			assert.Len(t, res, tt.expected)
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty slice", []int{}, 0},
		{"slice with 1 value", []int{1}, 1},
		{"correct slice", []int{4, 2, 1, 17, 21, 5}, 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maximum(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty slice", []int{}, 0},
		{"slice with 1 value", []int{42}, 42},
		{"correct slice", []int{4, 2, 1, 17, 21, 5}, 21},
		{"slice smaller than chunks", []int{9, 15, 3}, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxChunks(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}
