package helper

// ConvertToJSON convert map to json format expected by store procedure
func ConvertToJSON(products map[int]int) []map[string]any {
	var result []map[string]any
	for productId, quantity := range products {
		result = append(result, map[string]any{"productId": productId, "quantity": quantity})
	}
	return result
}
