package main

import (
	"slices"
	"testing"
)

func check(t *testing.T, nums []int, target int, want []int) {
	t.Helper()
	got := TwoSum(nums, target)
	slices.Sort(got)
	if !slices.Equal(got, want) {
		t.Errorf("TwoSum(%v, %d) = %v, want %v", nums, target, got, want)
	}
}

func TestExample(t *testing.T)        { check(t, []int{2, 7, 11, 15}, 9, []int{0, 1}) }
func TestPairNotAtStart(t *testing.T) { check(t, []int{3, 2, 4}, 6, []int{1, 2}) }
func TestSameValueTwice(t *testing.T) { check(t, []int{3, 3}, 6, []int{0, 1}) }
func TestNegativeNumbers(t *testing.T) {
	check(t, []int{-1, -2, -3, -4, -5}, -8, []int{2, 4})
}
