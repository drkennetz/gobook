package positive

func FindFirstPositiveNotInList(arr []int) int {
	found := []bool{true}
	missing := make([]bool, len(arr))
	found = append(found, missing...)

	for _, v := range arr {
		if (0 < v) && (v <= len(arr)) {
			found[v] = true
		}
	}
	for i, j := range found {
		if !j {
			return i
		}
	}
	return len(found)
}
