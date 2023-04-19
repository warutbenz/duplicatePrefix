package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []string{"finnomena", "fint", "fintech"}
	arr2 := []string{"finnomena", "dev", "tech"}

	fmt.Println("test 1:" + duplicatePrefix(arr))
	fmt.Println("test 2:" + duplicatePrefix(arr2))
}

func duplicatePrefix(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	prefix := arr[0]
	//loop data
	for i := 1; i < len(arr); i++ {
		//loop find prefix
		for !strings.HasPrefix(arr[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}
