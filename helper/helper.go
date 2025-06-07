package helper

func CheckString(input, fallback string) string {
	if input != "" {
		return input
	}
	return fallback
}
