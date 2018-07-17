package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort1(t *testing.T) {
	// Arrange
	input := []int{10, -23, -15, 1, 7, 55, 22}

	// Act
	actual := BubbleSort(input)

	// Assert
	expected := []int{-23, -15, 1, 7, 10, 22, 55}
	assert.Equal(t, expected, actual)
}

func TestBubbleSort2(t *testing.T) {
	// Arrange
	input := []int{5, 2, 1, 4, 2}

	// Act
	actual := BubbleSort(input)

	// Assert
	expected := []int{1, 2, 2, 4, 5}
	assert.Equal(t, expected, actual)
}
