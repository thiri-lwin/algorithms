package main

// Return index of the target if it's found in the list. If not, return -1
// This function assumed that input array is sorted array.
func BinarySearch[T NumberAndString](list []T, target T) int {
	first := 0
	last := len(list) - 1

	for first <= last {
		mid := (last + first) / 2
		if list[mid] == target {
			return mid
		}
		if target < list[mid] {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}
	return -1
}
