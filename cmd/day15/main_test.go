package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	assert.Equal(t, 1, memoryGame(2020, 1, 3, 2))
	assert.Equal(t, 10, memoryGame(2020, 2, 1, 3))
	assert.Equal(t, 27, memoryGame(2020, 1, 2, 3))
	assert.Equal(t, 78, memoryGame(2020, 2, 3, 1))
	assert.Equal(t, 438, memoryGame(2020, 3, 2, 1))
	assert.Equal(t, 1836, memoryGame(2020, 3, 1, 2))
}

func TestPuzzle2(t *testing.T) {
	assert.Equal(t, 175594, memoryGame(30000000, 0, 3, 6))
	assert.Equal(t, 2578, memoryGame(30000000, 1, 3, 2))
	assert.Equal(t, 3544142, memoryGame(30000000, 2, 1, 3))
	assert.Equal(t, 261214, memoryGame(30000000, 1, 2, 3))
	assert.Equal(t, 6895259, memoryGame(30000000, 2, 3, 1))
	assert.Equal(t, 18, memoryGame(30000000, 3, 2, 1))
	assert.Equal(t, 362, memoryGame(30000000, 3, 1, 2))
}
