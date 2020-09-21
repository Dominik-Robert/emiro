package helper

func SearchForIP(ips []string, ip string) bool {
	for _, value := range ips {
		if value == ip {
			return true
		}
	}
	return false
}
