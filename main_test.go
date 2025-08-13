package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("0 size", func(t *testing.T) {
		res := generateRandomElements(0)
		assert.Len(t, res, 0)
	})

	t.Run("correct size", func(t *testing.T) {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		size := rnd.Intn(1_000_000) + 1
		res := generateRandomElements(int(size))
		assert.Len(t, res, int(size))
	})
}

func TestMaximum(t *testing.T) {
	t.Run("empty slise", func(t *testing.T) {
		slice := []int{}
		res := maximum(slice)
		assert.Equal(t, -1, res)
	})

	t.Run("slise with 1 value", func(t *testing.T) {
		slice := []int{1}
		res := maximum(slice)
		assert.Equal(t, slice[0], res)
	})

	t.Run("correct slice", func(t *testing.T) {
		slice := []int{4, 2, 1, 17, 21, 5}
		res := maximum(slice)
		assert.Equal(t, 21, res)
	})
}

func TestMaxChunks(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		res := maxChunks(slice)
		assert.Equal(t, -1, res)
	})

	t.Run("slice with 1 value", func(t *testing.T) {
		slice := []int{42}
		res := maxChunks(slice)
		assert.Equal(t, 42, res)
	})

	t.Run("correct slice", func(t *testing.T) {
		slice := []int{4, 2, 1, 17, 21, 5}
		res := maxChunks(slice)
		assert.Equal(t, 21, res)
	})

	t.Run("slice smaller than chunks", func(t *testing.T) {
		slice := []int{9, 15, 3}
		res := maxChunks(slice)
		assert.Equal(t, 15, res)
	})
}
