package kubernetes

import (
	"fmt"
	"testing"
)

func TestLoadConfigDefaults(t *testing.T) {
	x, _ := run()
	fmt.Println(x)
}
