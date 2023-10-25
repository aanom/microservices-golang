package utils

// IsStringNilOrEmpty is utility function to check if string is nil or empty
func IsStringNilOrEmpty(sptr *string) bool {
	if sptr != nil {
		if *sptr != "" {
			return false
		}
	}

	return true
}
