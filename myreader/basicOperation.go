package myreader

func Contains(set []string, input string) bool {
	for _, value := range set {
		if input == value {
			return true
		}
	}
	return false
}
