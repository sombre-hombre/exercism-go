package flatten

// Flatten takes a nested list and returns a single flattened list
// with all values except nil/null
func Flatten(list interface{}) []interface{} {
	result := make([]interface{}, 0)

	array, ok := list.([]interface{})
	if !ok {
		if list != nil {
			result = append(result, list)
		}
		return result
	}

	for _, element := range array {
		switch el := element.(type) {
		case []interface{}:
			result = append(result, Flatten(el)...)
		case nil:
			continue
		default:
			result = append(result, el)
		}
	}

	return result
}
