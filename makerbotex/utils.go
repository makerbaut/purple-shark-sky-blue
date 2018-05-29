package main

import (
	"sort"
	"strconv"
)

// strMinusCharAtIndex returns a string with the character removed at the given index
func strMinusCharAtIndex(val string, index int) string {
	part1 := val[0:index]
	part2 := val[index+1 : len(val)]

	var res string
	for _, v := range part1 {
		res += string(v)
	}
	for _, v := range part2 {
		res += string(v)
	}

	return res
}

// permuations returns all the permutations of a given string
func permuations(prefix, val string, accum *[]string) {
	if len(val) == 0 {
		*accum = append(*accum, prefix)
	} else {
		for i := 0; i < len(val); i++ {
			permuations(prefix+string(val[i]), strMinusCharAtIndex(val, i), accum)
		}
	}
}

// getNumericalAnagrams returns a map of all the numerical anagrams so that they can be checked
// against with O(1) complexity
func getNumericalAnagrams(num string) map[string]bool {
	var perms []string
	permuations("", num, &perms)

	res := make(map[string]bool)
	for _, a := range perms {
		if a == num {
			// Skip the number itself since we don't want to compare against the
			// number itself, which will always yield true
			continue
		}

		res[a] = true
	}

	return res
}

// getMaxMinDifference will return the difference between the max and min
// elements of a given array
func getMaxMinDifference(nums []string) int {
	sort.Slice(nums, func(i, j int) bool {
		lhs, _ := strconv.Atoi(nums[i])
		rhs, _ := strconv.Atoi(nums[j])
		return lhs < rhs
	})

	min, _ := strconv.Atoi(nums[0])
	max, _ := strconv.Atoi(nums[len(nums)-1])
	return max - min
}

// hasAnagram returns true if the list of numbers contain numerical
// anagrams of each other
func hasAnagram(list []string) bool {
	for _, num := range list {
		// Skip single digit numbers
		if len(num) == 1 {
			continue
		}

		ags := getNumericalAnagrams(num)
		for _, l := range list {
			_, ok := ags[l]
			if ok {
				return true
			}
		}
	}

	return false
}

// hasTwoNumbersDividesTo177 returns true if any two numbers in the list
// divide to 177
func hasTwoNumbersDividesTo177(list []string) bool {
	for _, l := range list {
		n1, _ := strconv.Atoi(l)
		for _, k := range list {
			n2, _ := strconv.Atoi(k)
			// Check both n1/n2 and n2/n1 if n2 > n1 to prevent going through
			// the entire loop, since any two numbers can result in 177 when
			// divided.
			if n1/n2 == 177 && n1%n2 == 0 || n2/n1 == 177 && n2%n1 == 0 {
				return true
			}
		}
	}

	return false
}

// getChecksum is a utility function which will return the checksum given
// lists of list of numbers. This function is used in the unit-tests to test
// the other functions
func getChecksum(input [][]string) int {
	var checksum int
	for _, row := range input {
		if hasAnagram(row) || hasTwoNumbersDividesTo177(row) {
			continue
		}
		checksum += getMaxMinDifference(row)
	}

	return checksum
}
