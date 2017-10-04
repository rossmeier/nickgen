package nickgen

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(GenerateWord(i))
	}
}
