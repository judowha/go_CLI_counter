package counter

import (
	"fmt"
	"testing"
)

func TestCountFreq(t *testing.T) {
	result := CountFreq("hi this is GO CLI counter tool Here is a test for multiple file inputs", true)
	fmt.Print(result)
}
