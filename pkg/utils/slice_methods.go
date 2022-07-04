package utils

func IsStringInSlice(inputString string, slice []string) bool {
	for _, stringInSlice := range slice {
		if stringInSlice == inputString {
			return true
		}
	}

	return false
}
