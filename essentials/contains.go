package essentials

// Contains checks if elem exists in slice
func Contains(slice []string, elem string) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}
	return false
}
