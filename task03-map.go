package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	ids := make([]int, 0, len(input))

	for id := range input {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	result := make([]string, len(ids)-1)
	for _, value := range ids {
		result = append(result, value)
	}
	return
}
