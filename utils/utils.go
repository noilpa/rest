package utils


// check for an element in a slice
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	return false
}

