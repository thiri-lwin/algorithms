package main

// Return true if the target is found in the list. If not, return false
// This function assumed that input array is sorted array.
func RecersiveBinarySearch[T NumberAndString](list []T, target T) bool {
	if len(list) == 0 {
		return false
	}
	first := 0
	last := len(list) - 1
	mid := (first + last) / 2
	if list[mid] == target {
		return true
	}

	if target < list[mid] {
		return RecersiveBinarySearch(list[:mid], target)
	} else {
		return RecersiveBinarySearch(list[mid+1:], target)
	}
}
