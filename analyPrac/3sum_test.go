package analyPrac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test3sum1(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}

	want := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	got := threeSum(input)
	assert.Equal(t, want, got)
}

func Test3sum2(t *testing.T) {
	input := []int{}

	want := [][]int{}
	got := threeSum(input)
	assert.Equal(t, want, got)
}

func Test3sum3(t *testing.T) {
	input := []int{0}

	want := [][]int{}
	got := threeSum(input)
	assert.Equal(t, want, got)
}

func Test3sum4(t *testing.T) {
	input := []int{0, 0, 0, 0}

	want := [][]int{{0, 0, 0}}
	got := threeSum(input)
	assert.Equal(t, want, got)
}

func Test3sumLarge0Input(t *testing.T) {
	input := make([]int, 10000)

	want := [][]int{{0, 0, 0}}
	got := threeSum(input)
	assert.Equal(t, want, got)
}
