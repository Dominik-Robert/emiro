package helper

import "encoding/json"

func SearchForIP(ips []string, ip string) bool {
	for _, value := range ips {
		if value == ip {
			return true
		}
	}
	return false
}

func PrettyPrint(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}
