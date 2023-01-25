package main

// Return index of the target if it's found in the list. If not, return -1
func LinearSearch[T NumberAndString](list []T, target T) int {
	for i, element := range list {
		if element == target {
			return i
		}
	}
	return -1
}
