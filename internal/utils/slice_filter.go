package utils

var hardcodedItems = toMap([]string{
	// use hashmap as hashset to create O(1) filter
	"dog", "cat", "rabbit",
})

func SliceFilter(src []string) []string {
	if src == nil {
		// same behaviour like append - return slice on nil input
		return []string{}
	}
	result := make([]string, 0, len(src))

	for _, el := range src {
		if _, found := hardcodedItems[el]; !found {
			result = append(result, el)
		}
	}

	return result
}

// toMap converts slice of strings to map with string key. Go has no type hashset, use hashmap keyset
func toMap(strings []string) map[string]bool {
	result := make(map[string]bool)
	for _, s := range strings {
		result[s] = true
	}
	return result
}
