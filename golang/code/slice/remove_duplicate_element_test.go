package test

import (
	"fmt"
	"testing"
)

func removeDuplicateElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func TestRDE(t *testing.T) {
	s := []string{"hello", "world", "hello", "golang", "hello", "ruby", "php", "java"}
	fmt.Println(removeDuplicateElement(s))
}
