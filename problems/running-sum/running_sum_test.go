package main

import (
	"slices"
	"testing"
)

func check(t *testing.T, nums, want []int) {
	t.Helper()
	if got := RunningSum(nums); !slices.Equal(got, want) {
		t.Errorf("RunningSum(%v) = %v, want %v", nums, got, want)
	}
}

func TestExample(t *testing.T)   { check(t, []int{1, 2, 3, 4}, []int{1, 3, 6, 10}) }
func TestAllOnes(t *testing.T)   { check(t, []int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}) }
func TestMixed(t *testing.T)     { check(t, []int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}) }
func TestNegatives(t *testing.T) { check(t, []int{-1, 1, -1, 1}, []int{-1, 0, -1, 0}) }
