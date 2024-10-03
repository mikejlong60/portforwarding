package kubernetes

import (
	"fmt"
	"testing"
)

func TestGetIps(t *testing.T) {
	pods, _ := getIps()
	if len(pods) != 5 {
		t.Errorf("Expected 5 pods. Actual number of pods:%v\n", len(pods))
	}
	fmt.Println(pods)
}
