package main

import "testing"

func check(t *testing.T, nums []int, want bool) {
	t.Helper()
	if got := ContainsDuplicate(nums); got != want {
		t.Errorf("ContainsDuplicate(%v) = %v, want %v", nums, got, want)
	}
}

func TestHasDuplicate(t *testing.T)   { check(t, []int{1, 2, 3, 1}, true) }
func TestAllDistinct(t *testing.T)    { check(t, []int{1, 2, 3, 4}, false) }
func TestManyDuplicates(t *testing.T) { check(t, []int{1, 1, 1, 3, 3, 4, 3, 2}, true) }
func TestSingleElement(t *testing.T)  { check(t, []int{7}, false) }
