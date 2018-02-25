package flatten

// Flatten takes a nested list and returns a single flattened list
// with all values except nil/null
func Flatten(x interface{}) []interface{} {
	result := make([]interface{}, 0)

	switch x.(type) {
	case []interface{}:
		for _, el := range x.([]interface{}) {
			result = append(result, Flatten(el)...)
		}
	case nil:
		return result
	default:
		result = append(result, x)
	}

	return result
}
