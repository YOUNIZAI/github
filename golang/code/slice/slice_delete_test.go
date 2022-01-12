package test

import (
	"fmt"
	"testing"
)
func TestMy(t *testing.T) {
	b :=[]int{1,2,3,4,5}
	i := 0
	for k, n := range b {
		if k==2 || k==4 {
			b[i] = n
			i++
		}
	}
	b = b[:i]
  fmt.Printf("b:%v \n",b)
}
//test result:
//=== RUN   TestMy
//b:[3 5]
