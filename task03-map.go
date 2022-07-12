package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	keyMap := []int{}

	for i := range input {
		keyMap = append(keyMap, i)
	}
	sort.Ints(keyMap)

	for i := 0; i < len(keyMap); i++ {
		result = append(result, input[keyMap[i]])
	}
	return
}
