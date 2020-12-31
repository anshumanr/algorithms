package analyPrac

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	l := len(nums)

	if l == 0 || nums[0] > 0 || nums[l-1] < 0 {
		return [][]int{}
	}

	for i := 0; i < l-2; i++ {
		j := i + 1
		k := l - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				z := []int{nums[i], nums[j], nums[k]}
				if isUnique(res, z) {
					res = append(res, z)
				}
				j++
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return res
}

func isUnique(sum [][]int, c []int) bool {
	l := len(sum)

	for i := 0; i < l; i++ {
		if isEqual(sum[i], c) {
			return false
		}
	}

	return true
}

func isEqual(a, b []int) bool {
	l1 := len(a)
	l2 := len(b)

	if l1 != l2 {
		return false
	}

	for i := 0; i < l1; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
