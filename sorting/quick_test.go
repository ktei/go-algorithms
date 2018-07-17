package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort1(t *testing.T) {
	// Arrange
	input := []int{10, -23, -15, 1, 7, 55, 22}

	// Act
	actual := QuickSort(input, 0, 0, len(input)-1)

	// Assert
	expected := []int{-23, -15, 1, 7, 10, 22, 55}
	assert.Equal(t, expected, actual)
}

func TestQuickSort2(t *testing.T) {
	// Arrange
	input := []int{5, 2, 1, 4, 2}

	// Act
	actual := QuickSort(input, 0, 0, len(input)-1)

	// Assert
	expected := []int{1, 2, 2, 4, 5}
	assert.Equal(t, expected, actual)
}

func TestQuickSort3(t *testing.T) {
	// Arrange
	input := []int{99, 2, 1, 4, 2, 10, -5, -66}

	// Act
	actual := QuickSort(input, 0, 0, len(input)-1)

	// Assert
	expected := []int{-66, -5, 1, 2, 2, 4, 10, 99}
	assert.Equal(t, expected, actual)
}
