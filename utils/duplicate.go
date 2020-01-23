package utils

// CheckForDuplicateIP finds duplicate IP in 2 arrays
func CheckForDuplicateIP(ips []string) bool {
	var isDuplicate bool
	ipMap := make(map[string]bool)

	for _, ip := range ips {
		if _, exists := ipMap[ip]; exists {
			isDuplicate = true
			break
		}
		ipMap[ip] = true
	}

	return isDuplicate
}
