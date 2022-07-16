package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	ids := make([]int, 0, len(input))

	for id := range input {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	for _, valu := range input {
		result = append(result, valu)
	}
	return
}
