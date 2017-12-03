package essentials

// Choice offer a one-liner func to choose
func Choice(chooser bool, param1, param2 interface{}) interface{} {
	if chooser {
		return param1
	} else {
		return param2
	}
}
